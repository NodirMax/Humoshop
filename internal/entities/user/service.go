package user

func GetUserService(u UserModel) (user UserModel, err error) {
	user, err = GetUserFromDB(u)
	if err != nil {
		return
	}
	return
}

func CreateUserService(user UserModel) (err error) {
	err = CreateUserDB(user)
	if err != nil {
		return
	}
	return
}

func UpdateUserService(user UserModel) (err error) {
	err = UpdateUserDB(user)
	if err != nil {
		return
	}
	return
}

func DeleteUserService(user UserModel) (err error) {
	err = DeleteUserDB(user)
	if err != nil {
		return
	}
	return
}
