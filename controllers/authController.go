package controllers

import (
	"database/sql"
	"fmt"

	"api_go_store/jwt"

	"net/http"

	"api_go_store/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func PasswordHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func PasswordCek(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func Register(c *gin.Context) {
	user := models.User{}
	// ambil dari dari json
	if err := c.Bind(&user); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	// perintah buka koneksi database
	db := c.MustGet("db").(*sql.DB)
	pass, _ := PasswordHash(user.Password)
	// ekseskusi perintah ke dalam bentuk sql
	if data, _ := db.Query("INSERT INTO `users`(`nama`, `username`, `password`) VALUES (?,?,?)", user.Nama, user.Username, pass); data != nil {
		c.JSON(http.StatusOK, gin.H{"data": user})
	}
}

type Input struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Logins(c *gin.Context) {
	// deklarasi type data yang akan diambil
	var input Input
	// deklarasi type data dari model
	var user models.User
	// ambil data dari json dan cek jika data error atau tidak
	if err := c.Bind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	// buka koneksi ke database
	db := c.MustGet("db").(*sql.DB)
	// cek apakah username ada
	db.QueryRow("SELECT id, password FROM users WHERE username = ?", input.Username).Scan(&user.Id, &user.Password)
	// cek apakah password betul atau salah
	hasil := PasswordCek(input.Password, user.Password)
	// cek jika password benar
	if hasil == nil {
		// generate token
		token, err := jwt.GenerateToken(user.Id)
		// return token via json
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"Error": "Error"})
		} else {
			c.JSON(http.StatusOK, gin.H{"token": token})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{"Error": "Password salah"})
	}
}

func UserId(c *gin.Context) {
	// panggil model user
	var userDetail models.UserDetail
	// cek token yang dikirim dari header
	id, err := jwt.ExtractTokenID(c)
	// jika token yg di kirim salah return pesan
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	// buka koneksi ke database
	db := c.MustGet("db").(*sql.DB)
	// cek data user berdasarkan id dari token
	data := db.QueryRow("SELECT id,nama,username,gambar FROM users WHERE id = ?", id).Scan(&userDetail.Id, &userDetail.Nama, &userDetail.Username, &userDetail.Gambar)
	fmt.Println(data)
	if data != nil {
		// return pesan error
		c.JSON(http.StatusInternalServerError, "Not Found")
	} else {
		// return data json berdasarkan data user
		c.JSON(http.StatusOK, userDetail)
	}
}
