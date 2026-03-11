package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/viktare/go-commerce/internal/repository"
)

type CreateProductDTO struct {
	Title string `json:"title" binding:"required"`
	Code  string `json:"code" binding:"required"`
}

func CreateProductHandler(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var input CreateProductDTO

		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		product, err := repository.CreateProduct(pool, input.Title, input.Code)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		ctx.JSON(http.StatusCreated, product)
	}
}
