package main

import (
	"database/sql"
	"errors"
)

const insertProduct = `
		INSERT INTO product (name)
		VALUES ($1)
		RETURNING id`

func InsertProduct(db *sql.DB, dto ProductDto) (product Product, err error) {
	var id int
	err = db.QueryRow(insertProduct, dto.Name).Scan(&id)
	if err != nil {
		return
	}
	return GetProductByID(db, id)
}

const updateProduct = `
	UPDATE 
	    product
	SET
		"name" = $2
	WHERE id = $1;
 `

func UpdateByID(db *sql.DB, productId int, dto ProductDto) (product Product, err error) {
	_, err = GetProductByID(db, productId)
	if err != nil {
		return
	}
	_, err = db.Exec(updateProduct, productId, dto.Name)
	if err != nil {
		return
	}
	return GetProductByID(db, productId)

}

const selectProductByID = `
	SELECT
		id ,
		"name"
	FROM product
	WHERE id = $1`

func GetProductByID(db *sql.DB, productId int) (product Product, err error) {
	result, err := db.Query(selectProductByID, productId)
	if err != nil {
		return
	}
	if !result.Next() {
		err = errors.New("invalid ID ")
		return
	}
	err = result.Scan(&product.ID, &product.Name)
	if err != nil {
		return
	}
	return
}

const deleteProduct = `
	DELETE FROM
		"product"
	WHERE
		id = $1;
`

func DeleteByID(db *sql.DB, productId int) (err error) {
	_, err = GetProductByID(db, productId)
	if err != nil {
		return
	}
	_, err = db.Exec(deleteProduct, productId)
	if err != nil {
		return
	}
	return
}

const selectProducts = `
	SELECT
		id,
		"name"				
	FROM "product"
 `

func SelectProducts(db *sql.DB) (products []Product, err error) {
	products = make([]Product, 0)

	results, err := db.Query(selectProducts)
	if err != nil {
		return products, err // id does not exist
	}

	for results.Next() {
		product := Product{}
		err = results.Scan(
			&product.ID,
			&product.Name,
		)
		if err != nil {
			return
		}
		products = append(products, product)
	}
	return
}
