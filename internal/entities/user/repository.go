package user

import (
	"humoshop/pkg/db"
	"log"
)

// Запрос к базе данных о пользователье
func GetUserFromDB(user UserModel) (u UserModel, err error) {
	row := db.DB.QueryRow("SELECT * FROM users WHERE id=$1", user.Id)

	err = row.Scan(&u.Id, &u.Name, &u.Surname, &u.Phone, &u.Login, &u.Password, &u.DataTIME)
	if err != nil {
		log.Println("Ошибка row")
	}
	return
}

// Создание нового пользователь
func CreateUserDB(user UserModel) (err error) {
	_, err = db.DB.Exec("INSERT INTO users(name, surname, phone, login, password) VALUES($1, $2, $3, $4, $5)", user.Name, user.Surname, user.Phone, user.Login, user.Password)
	return err
}

// Обновление данных пользователья в Базе Данных
func UpdateUserDB(user UserModel) (err error) {
	_, err = db.DB.Exec("UPDATE users SET name=$1, surname=$2, phone=$3, login=$4, password=$5 WHERE id=$6", user.Name, user.Surname, user.Phone, user.Login, user.Password, user.Id)
	return err
}

// Удаление данных пользователья в Базе Данных
func DeleteUserDB(user UserModel) (err error) {
	_, err = db.DB.Exec("DELETE FROM users WHERE id=$1", user.Id)
	return err
}
