package main

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func SelectProductByID(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("product_id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			err = errors.New("can not convert to int ")
			c.AbortWithStatusJSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		product, err := GetProductByID(db, idInt)
		if err != nil {
			c.AbortWithStatusJSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(200, product)

	}
}

func AddProduct(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		dto := ProductDto{}
		err := c.ShouldBindJSON(&dto)
		if err != nil {
			c.AbortWithStatusJSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		product, err := InsertProduct(db, dto)
		if err != nil {
			c.AbortWithStatusJSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(200, product)
	}
}

func UpdateProductByID(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("product_id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			err = errors.New("can not convert to int ")
			c.AbortWithStatusJSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		dto := ProductDto{}
		err = c.ShouldBindJSON(&dto)
		if err != nil {
			c.AbortWithStatusJSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		product, err := UpdateByID(db, idInt, dto)
		if err != nil {
			c.AbortWithStatusJSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(200, product)
	}
}

func DeleteProductByID(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("product_id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			err = errors.New("can not convert to int ")
			c.AbortWithStatusJSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		err = DeleteByID(db, idInt)
		if err != nil {
			c.AbortWithStatusJSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"message": fmt.Sprintf("product with id=%v deleted successfully", idInt),
		})
	}
}

func GetProducts(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		products, err := SelectProducts(db)
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(200, products)
		return
	}

}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
