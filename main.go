package main

import (
	"api_go_store/config"
	"api_go_store/router"
	"fmt"
)

func main() {
	fmt.Println("Hallo")

	db := config.Koneksi()
	defer db.Close()

	r := router.Router(db)
	r.Run()
}
