package handlers

import (
	"encoding/json"
	"humoshop/internal/entities/category"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	var i int64
	for i = 1; i <= 5; i ++{
		z := category.CategoryStruct{Id: i}
		result, err := category.CategoryGETDB(z)
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("Ошибка при получение данных из базы данных"))
			return
		}

		resp, _ := json.Marshal(result)
		w.WriteHeader(http.StatusOK)
		w.Write(resp)
		
	}
	return
}
