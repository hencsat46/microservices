package models

type User struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
	Name     string `json:"Name"`
	Surname  string `json:"Surname"`
}
