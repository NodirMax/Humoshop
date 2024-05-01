package middleware

import (
	"encoding/json"
	"humoshop/internal/entities/user"
	"net/http"
)

func UserChecking(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var user user.UserModel
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			w.WriteHeader(401)
			w.Write([]byte("Ошибка при отправке данных"))
		}
		
		// Проверка наличии Имени пользователья
		if user.Name == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Номер телефона не введён!"))
			return
		}

		// Проверка наличие номеров телефона
		if user.Phone == ""{
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Номер телефона не введён!"))
			return
		}

		// Проверка наличии пароля
		if user.Password == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Номер телефона не введён!"))
			return
		}

		// Проверка наличии логина!
		if user.Login == ""{
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Номер телефона не введён!"))
			return
		}
	
	})
}