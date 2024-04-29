package category

import (
	"humoshop/internal/entities/product"
	"humoshop/pkg/db"
	"log"
)

// Получение данных о категории из БД
func CategoryGETDB(category CategoryStruct) (cat []product.ProductDatabaseStruct, err error) {
	rows, err := db.DB.Query("SELECT * FROM product WHERE category_id=$1", category.Id)
	if err != nil {
		return
	}
	defer rows.Close()

	var c product.ProductDatabaseStruct
	for rows.Next() {
		err := rows.Scan(&c.Product_id, &c.Name, &c.Price, &c.In_stock, &c.Category_id)
		if err != nil {
			log.Println("Ошибка row")
			continue
		}
		cat = append(cat, c)
	}
	return
}

// Создание новой категории в БД
func CategoryCreateDB(cat CategoryStruct) (err error) {
	_, err = db.DB.Exec("INSERT INTO category(category_name) VALUES($1)", cat.CategoryName)
	return err
}

// Обновление данных категории в Базе Данных
func UpdateCategoryDB(cat CategoryStruct) (err error) {
	_, err = db.DB.Exec("UPDATE category SET category_name=$1 WHERE id=$2", cat.CategoryName, cat.Id)
	return err
}

// Удаление данных о категории в Базе Данных
func DeleteCategoryDB(cat CategoryStruct) (err error) {
	_, err = db.DB.Exec("DELETE FROM category WHERE id=$1", cat.Id)
	return err
}
