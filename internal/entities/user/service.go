package user

func GetUserService(login string) (users []UserDatabase, err error) {
	users, err = GetUserFromDB(login)
	if err != nil {
		return
	}
	return
}

func CreateUserService(user CreateUserStruct) (err error) {
	err = CreateUserDB(user)
	if err != nil {
		return
	}
	return
}

func UpdateUserService(user UpdateUserStruct) (err error) {
	err = UpdateUserDB(user)
	if err != nil {
		return
	}
	return
}

func DeleteUserService(user DeleteUserStruct) (err error) {
	err = DeleteUserDB(user)
	if err != nil {
		return
	}
	return
}
