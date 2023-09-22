package product

import (
	"database/sql"
	"fmt"
	"go-api/entity/model"
	"log"
)

func GetProductName(productID int, db *sql.DB) (model.Product, error) {
	var product model.Product

	// create the select sql query
	sqlStatement := `SELECT * FROM product WHERE product_id=$1`

	// execute the sql statement
	row := db.QueryRow(sqlStatement, productID)

	err := row.Scan(&product.ProductID, &product.Name)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return product, nil
	case nil:
		return product, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	return product, err
}
