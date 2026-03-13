package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/viktare/go-commerce/internal/config"
	"github.com/viktare/go-commerce/internal/database"
	"github.com/viktare/go-commerce/internal/handlers"
)

func main() {
	cfg, err := config.Load()

	if err != nil {
		log.Fatal("Failed to load configuration:", err)
	}

	pool, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	defer pool.Close()

	router := gin.Default()
	router.SetTrustedProxies(nil)
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Server running",
		})
	})

	router.GET("/products", handlers.GetProductsHandler(pool))
	router.POST("/products", handlers.CreateProductHandler(pool))

	router.Run(":" + cfg.Port)
}
