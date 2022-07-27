package controllers

import (
	"database/sql"
	"net/http"

	"api_go_store/models"

	"github.com/gin-gonic/gin"
)

func GetProduk(c *gin.Context) {
	db := c.MustGet("db").(*sql.DB)
	var produk []models.Produk
	var list models.Produk

	data, err := db.Query("SELECT * FROM produks")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	for data.Next() {
		var err = data.Scan(&list.ID, &list.Nama_produk, &list.Deskripsi, &list.Harga, &list.Stok, &list.Gambar, &list.Raiting)

		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
		}

		produk = append(produk, list)
	}

	c.JSON(http.StatusOK, produk)
}
