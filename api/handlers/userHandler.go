package handlers

import (
	"encoding/json"
	"humoshop/internal/entities/user"
	"net/http"
)

// получение данных от пользователья
func GetUser(w http.ResponseWriter, r *http.Request) {
	var User user.UserModel
	err := json.NewDecoder(r.Body).Decode(&User)
	if err != nil {
		w.WriteHeader(401)
		w.Write([]byte("Ошибка при отправке данных"))
		return
	}

	res, err := user.GetUserService(User)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Ошибка при получение данных из базы данных"))
	}
	resp, _ := json.Marshal(res)

	w.WriteHeader(http.StatusOK)
	w.Write(resp)
	return
}

// Регистрация пользователья
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var NewUser user.UserModel

	err := json.NewDecoder(r.Body).Decode(&NewUser)
	if err != nil {
		w.WriteHeader(401)
		w.Write([]byte("Ошибка при отправке данных"))
		return
	}

	err = user.CreateUserService(NewUser)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Ошибка сервера при сохранение"))
		return
	}

	w.WriteHeader(200)
	w.Write([]byte("Успешно зарегестрирован!"))
	return
}

// изменение данных пользователья
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var UpdateUser user.UserModel

	err := json.NewDecoder(r.Body).Decode(&UpdateUser)
	if err != nil {
		w.WriteHeader(401)
		w.Write([]byte("Ошибка при отправке данных"))
		return
	}

	err = user.UpdateUserService(UpdateUser)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Ошибка сервера при обновление данных"))
		return
	}

	w.WriteHeader(200)
	w.Write([]byte("Данные успешно обновлени!"))
	return
}

// удаление пользователья
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	var UserID user.UserModel

	err := json.NewDecoder(r.Body).Decode(&UserID)
	if err != nil {
		w.WriteHeader(401)
		w.Write([]byte("Ошибка при отправке данных"))
		return
	}

	err = user.DeleteUserService(UserID)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Ошибка при удаление данных!"))
	}

	w.WriteHeader(200)
	w.Write([]byte("Данные пользователья удалены успешно!"))
	return
}
