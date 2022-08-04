package router

import (
	"api_go_store/controllers"
	"api_go_store/middleware"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func Router(db *sql.DB) *gin.Engine {
	r := gin.Default()
	r.Use(func(ctx *gin.Context) {
		ctx.Set("db", db)
	})

	belumlogin := r.Group("/api")
	{
		belumlogin.Use(middleware.MiddlewareRoute())
		belumlogin.POST("/register", controllers.Register)
		belumlogin.POST("/login", controllers.Logins)
		belumlogin.GET("/getProduk", controllers.GetProduk)
	}

	sudahlogin := r.Group("/auth")
	{
		sudahlogin.Use(middleware.MiddlewareToken())
		sudahlogin.GET("/user", controllers.UserId)
		sudahlogin.POST("/addCart", controllers.AddCart)
		sudahlogin.GET("/getCart/:id", controllers.GetCart)
		sudahlogin.POST("/CartAksi", controllers.CartAksi)
	}

	return r
}
