package handlers

import (
	"encoding/json"
	"humoshop/internal/entities/category"
	"net/http"
)

// func CategoryHandler(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	categoryName := vars["category_name"]
// 	fmt.Fprint(w, categoryName)
// }

//Получение данных о категории
func CategoryGET(w http.ResponseWriter, r *http.Request) {
	var CategoryGet category.CategoryStruct

	err := json.NewDecoder(r.Body).Decode(&CategoryGet)
	if err != nil {
		w.WriteHeader(401)
		w.Write([]byte("Ошибка при отправке данных"))
		return
	}

	result, err := category.CategoryGETDB(CategoryGet)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Ошибка при получение данных из базы данных"))
		return
	}

	resp, _ := json.Marshal(result)
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
	return
}

// Добавление новой категории
func CategoryCreate(w http.ResponseWriter, r *http.Request) {
	var CreateCategory category.CategoryStruct

	err := json.NewDecoder(r.Body).Decode(&CreateCategory)
	if err != nil {
		w.WriteHeader(401)
		w.Write([]byte("Ошибка при отправке данных"))
		return
	}

	err = category.CategoryCreateDB(CreateCategory)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Ошибка сервера при сохранение"))
		return
	}

	w.WriteHeader(200)
	w.Write([]byte("Категория успешно добавлен"))
	return
}

// Обновление данных о категории
func CategoryUpdate(w http.ResponseWriter, r *http.Request) {
	var UpdateCategory category.CategoryStruct

	err := json.NewDecoder(r.Body).Decode(&UpdateCategory)
	if err != nil {
		w.WriteHeader(401)
		w.Write([]byte("Ошибка при отправке данных"))
		return
	}

	err = category.UpdateCategoryDB(UpdateCategory)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Ошибка сервера при Обновление данных"))
		return
	}

	w.WriteHeader(200)
	w.Write([]byte("Данные о категории успешно изменени"))
	return
}

// Удаления категории
func CategoryDelete(w http.ResponseWriter, r *http.Request) {
	var DeleteCategory category.CategoryStruct
	err := json.NewDecoder(r.Body).Decode(&DeleteCategory)
	if err != nil {
		w.WriteHeader(401)
		w.Write([]byte("Ошибка при отправке данных"))
		return
	}

	err = category.DeleteCategoryDB(DeleteCategory)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Ошибка сервера при удаленые данных"))
		return
	}

	w.WriteHeader(200)
	w.Write([]byte("Данные о продукте успешно удалены"))
	return
}
