package controllers

import (
	"api_go_store/models"
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddCart(c *gin.Context) {
	var inputCart models.Cart
	if err := c.Bind(&inputCart); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	db := c.MustGet("db").(*sql.DB)
	data, err := db.Query("INSERT INTO `carts`(`id_produk`, `id_user`, `jumlah`, `tanggal`) VALUES (?,?,?,?)", inputCart.Id_produk, inputCart.Id_user, inputCart.Jumlah, inputCart.Tanggal)

	if err != nil {
		c.JSON(http.StatusOK, data)
	}

}

func GetCart(c *gin.Context) {
	id_user := c.Param("id")
	var cartdetail []models.CartDetail
	var cartList models.CartDetail
	db := c.MustGet("db").(*sql.DB)

	data, err := db.Query("SELECT carts.id, produks.id, carts.id_user, carts.jumlah, carts.tanggal, produks.nama_produk, produks.deskripsi, produks.harga, produks.stok, produks.gambar, produks.raiting FROM carts JOIN produks ON carts.id_produk = produks.id WHERE carts.id_user = ?", id_user)
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
