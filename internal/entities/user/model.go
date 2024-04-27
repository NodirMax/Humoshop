package user

type UserModel struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Phone    string `json:"phone"`
	Login    string `json:"login"`
	Password string `json:"password"`
	DataTIME string `json:"datatime"`
}
