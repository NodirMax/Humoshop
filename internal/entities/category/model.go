package category

type CategoryStruct struct {
	Id           int64
	CategoryName string
}

type CategoryGETStruct struct {
	Id int64 `json:"id"`
}
