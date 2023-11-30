package models

type RecordModel struct {
	RecordId int
	UserId   int
}

type User struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
	Name     string `json:"Name"`
	Surname  string `json:"Surname"`
	Id       int    `json:"Id"`
}

type Response struct {
	Status  int
	Payload interface{}
}
