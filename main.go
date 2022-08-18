package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "Pharmacies"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	engine := gin.Default()
	engine.GET("/ping", ping)

	productPath := engine.Group("/products")

	productPath.POST("", AddProduct(db))
	productPath.GET("/:product_id", SelectProductByID(db))
	productPath.PUT("/:product_id", UpdateProductByID(db))
	productPath.DELETE("/:product_id", DeleteProductByID(db))
	productPath.GET("", GetProducts(db))

	engine.Run()
}
