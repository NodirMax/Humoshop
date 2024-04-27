package handlers

import (
	"encoding/json"
	"humoshop/internal/entities/product"
	"net/http"
)

// Получение данных о продукте
func ProductGET(w http.ResponseWriter, r *http.Request) {
	var GetProduct product.ProductDatabaseStruct

	err := json.NewDecoder(r.Body).Decode(&GetProduct)
	if err != nil {
		w.WriteHeader(401)
		w.Write([]byte("Ошибка при отправке данных"))
		return
	}

	result, err := product.GetProductDB(GetProduct)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Ошибка при получение данных из базы данных"))
	}

	resp, _ := json.Marshal(result)
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
	return
}

// Добавление продукта
func ProductCreate(w http.ResponseWriter, r *http.Request) {
	var CreateProduct product.ProductDatabaseStruct
	err := json.NewDecoder(r.Body).Decode(&CreateProduct)
	if err != nil {
		w.WriteHeader(401)
		w.Write([]byte("Ошибка при отправке данных"))
		return
	}

	err = product.CreateProductDB(CreateProduct)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Ошибка сервера при сохранение"))
		return
	}

	w.WriteHeader(200)
	w.Write([]byte("Продукт успешно добавлен"))
	return
}

// Обновление данных продукта
func ProductUpdate(w http.ResponseWriter, r *http.Request) {
	var Product product.ProductDatabaseStruct

	err := json.NewDecoder(r.Body).Decode(&Product)
	if err != nil {
		w.WriteHeader(401)
		w.Write([]byte("Ошибка при отправке данных"))
		return
	}

	err = product.UpdateProductDB(Product)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Ошибка сервера при Обновление данных"))
		return
	}

	w.WriteHeader(200)
	w.Write([]byte("Данные о продукте изменени успешно!"))
	return
}

// Удаление данных о продукте
func ProductDelete(w http.ResponseWriter, r *http.Request) {
	var DeleteUser product.ProductDatabaseStruct
	err := json.NewDecoder(r.Body).Decode(&DeleteUser)
	if err != nil {
		w.WriteHeader(401)
		w.Write([]byte("Ошибка при отправке данных"))
		return
	}

	err = product.DeleteProductDB(DeleteUser)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Ошибка сервера при удаленые данных"))
		return
	}

	w.WriteHeader(200)
	w.Write([]byte("Данные о продукте успешно удалены"))
	return
}
