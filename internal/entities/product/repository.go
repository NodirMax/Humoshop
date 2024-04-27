package product

import (
	"humoshop/pkg/db"
	"log"
)

// Запрос к базе данных о продукте
func GetProductDB(product_id ProductDatabaseStruct) (product []ProductDatabaseStruct, err error) {
	rows, err := db.DB.Query("SELECT * FROM product WHERE product_id=$1", product_id.Product_id)
	if err != nil {
		return
	}
	defer rows.Close()

	var p ProductDatabaseStruct
	for rows.Next() {
		err := rows.Scan(&p.Product_id, &p.Name, &p.Price, &p.In_stock)
		if err != nil {
			log.Println("Ошибка row")
			continue
		}
	}
	product = append(product, p)
	return
}

// Добавление нового продукта В базу данных
func CreateProductDB(user ProductDatabaseStruct) (err error) {
	_, err = db.DB.Exec("INSERT INTO product(name, price, in_stock) VALUES ($1, $2, $3)", user.Name, user.Price, user.In_stock)
	return
}

// Обновление информации о продукте
func UpdateProductDB(p ProductDatabaseStruct) (err error) {
	_, err = db.DB.Exec("UPDATE product SET name=$1, price=$2, in_stock=$3 WHERE product_id=$4", p.Name, p.Price, p.In_stock, p.Product_id)
	return
}

// Удаление продукта из Базы Данных
func DeleteProductDB(p ProductDatabaseStruct) (err error) {
	_, err = db.DB.Exec("DELETE FROM product WHERE product_id=$1", p.Product_id)
	return
}
