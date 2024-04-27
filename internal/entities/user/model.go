package user

type UserDatabase struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Phone    string `json:"phone"`
	Login    string `json:"login"`
	Password string `json:"password"`
	DataTIME string `json:"datatime"`
}

type CreateUserStruct struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Phone    string `json:"phone"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type UpdateUserStruct struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Phone    string `json:"phone"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type DeleteUserStruct struct {
	Id int64 `json:"id"`
}
