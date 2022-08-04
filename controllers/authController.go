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
	var user models.User
	// ambil dari dari json
	if err := c.BindJSON(&user); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Println(user)
	// perintah buka koneksi database
	db := c.MustGet("db").(*sql.DB)
	pass, _ := PasswordHash(user.Password)
	// ekseskusi perintah ke dalam bentuk sql
	if data, _ := db.Query("INSERT INTO `user`(`nama`, `username`, `password`,`gambar`) VALUES (?,?,?,?)", user.Nama, user.Username, pass, ""); data != nil {
		c.JSON(http.StatusOK, data)
	} else {
		c.JSON(http.StatusOK, "Error")
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
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	// buka koneksi ke database
	db := c.MustGet("db").(*sql.DB)
	// cek apakah username ada
	sqlStatement := "SELECT id, password FROM user WHERE username = '" + input.Username + "'"
	db.QueryRow(sqlStatement).Scan(&user.Id, &user.Password)
	// cek apakah password betul atau salah
	err := PasswordCek(input.Password, user.Password)
	// cek jika password benar
	if err == nil {

		// generate token
		token, err := jwt.GenerateToken(user.Id)
		// set cookie
		c.SetCookie("token", token, 60*60*24, "/", "", true, true)
		result := map[string]interface{}{
			"token": token,
			"id":    user.Id,
		}

		// return token via json
		if err != nil {
			c.JSON(http.StatusBadRequest, "Error")
		} else {
			c.JSON(http.StatusOK, result)
		}
	} else {
		c.JSON(http.StatusBadRequest, "Password salah")
	}
}

func UserId(c *gin.Context) {
	// panggil model user
	var userDetail models.UserDetail
	// cek token yang dikirim dari header
	id, err := jwt.ExtractTokenID(c)
	// jika token yg di kirim salah return pesan
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	// buka koneksi ke database
	db := c.MustGet("db").(*sql.DB)
	// cek data user berdasarkan id dari token
	data := db.QueryRow("SELECT id,nama,username,gambar FROM user WHERE id = ?", id).Scan(&userDetail.Id, &userDetail.Nama, &userDetail.Username, &userDetail.Gambar)

	if data != nil {
		// return pesan error
		c.JSON(http.StatusInternalServerError, "Not Found")
	} else {
		// return data json berdasarkan data user
		c.JSON(http.StatusOK, userDetail)
	}
}
