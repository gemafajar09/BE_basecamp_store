package controllers

import (
	"api_go_store/models"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddCart(c *gin.Context) {
	var cart models.Cart
	var inputCart models.Cart
	if err := c.Bind(&inputCart); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	db := c.MustGet("db").(*sql.DB)
	cek := db.QueryRow("SELECT id, jumlah FROM cart WHERE id_produk = ? AND id_user = ?", inputCart.Id_produk, inputCart.Id_user).Scan(&cart.Id, &cart.Jumlah)
	if cek == nil {
		var jumlah int = cart.Jumlah + inputCart.Jumlah

		_, err := db.Query("UPDATE `cart` SET `jumlah`=? WHERE `id`=?", jumlah, cart.Id)
		fmt.Println(err)
		if err == nil {
			c.JSON(http.StatusOK, true)
		} else {
			c.JSON(http.StatusOK, false)
		}
	} else {
		_, err := db.Query("INSERT INTO `cart`(`id_produk`, `id_user`, `jumlah`, `tanggal`) VALUES (?,?,?,?)", inputCart.Id_produk, inputCart.Id_user, inputCart.Jumlah, inputCart.Tanggal)

		if err == nil {
			c.JSON(http.StatusOK, true)
		} else {
			c.JSON(http.StatusOK, false)
		}

	}
}

func CartAksi(c *gin.Context) {
	var cart models.Cart
	var inputCart models.MinOrMax
	if err := c.BindJSON(&inputCart); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	db := c.MustGet("db").(*sql.DB)
	cek := db.QueryRow("SELECT id, jumlah FROM cart WHERE id = ? ", inputCart.Id).Scan(&cart.Id, &cart.Jumlah)
	c.JSON(http.StatusOK, inputCart)
	if cek == nil {
		var jumlah int
		if inputCart.Kode == "+" {
			jumlah = cart.Jumlah + inputCart.Jumlah
		} else {
			jumlah = cart.Jumlah - inputCart.Jumlah
		}

		_, err := db.Query("UPDATE `cart` SET `jumlah`=? WHERE `id`=?", jumlah, cart.Id)
		fmt.Println(err)
		if err == nil {
			c.JSON(http.StatusOK, true)
		} else {
			c.JSON(http.StatusOK, false)
		}
	}
	c.JSON(http.StatusOK, false)
}

func GetCart(c *gin.Context) {
	id_user := c.Param("id")
	var cartdetail []models.CartDetail
	var cartList models.CartDetail
	db := c.MustGet("db").(*sql.DB)

	data, err := db.Query("SELECT cart.id, produk.id, cart.id_user, cart.jumlah, cart.tanggal, produk.nama_produk, produk.deskripsi, produk.harga, produk.stok, produk.gambar, produk.raiting FROM cart JOIN produk ON cart.id_produk = produk.id WHERE cart.id_user = ?", id_user)
	for data.Next() {
		var stok string
		err = data.Scan(
			&cartList.Id,
			&cartList.Id_produk,
			&cartList.Id_user,
			&cartList.Jumlah,
			&cartList.Tanggal,
			&cartList.Nama_produk,
			&cartList.Deskripsi,
			&cartList.Harga,
			&stok,
			&cartList.Gambar,
			&cartList.Raiting)

		// conversi int ke string
		cartList.Stok, _ = strconv.Atoi(stok)
		if err != nil {
			c.JSON(http.StatusBadGateway, err.Error())
		}

		cartdetail = append(cartdetail, cartList)
	}

	if err != nil {
		c.JSON(http.StatusNotFound, data)
	}
	c.JSON(http.StatusOK, cartdetail)

}
