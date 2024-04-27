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

func CategoryGET(w http.ResponseWriter, r *http.Request) {
	var CategoryGet category.CategoryGETStruct

	err := json.NewDecoder(r.Body).Decode(&CategoryGet)
	if err != nil {
		w.WriteHeader(401)
		w.Write([]byte("Ошибка при отправке данных"))
		return
	}

	// err = category.

}
