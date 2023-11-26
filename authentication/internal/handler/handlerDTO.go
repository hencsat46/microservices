package handler

type userDTO struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
	Name     string `json:"Name"`
	Surname  string `json:"Surname"`
	Id       int    `json:"Id"`
}
