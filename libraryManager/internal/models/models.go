package models

type RecordModel struct {
	RecordId int
	UserId   int
}

type Response struct {
	Status  int
	Payload interface{}
}
