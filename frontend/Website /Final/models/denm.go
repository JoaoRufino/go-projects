package models

type DENM struct {
	Accident int
	Id       string
	Content  string
}

func NewDENM(accident int, id, content string) *DENM {
	return &DENM{accident, id, content}
}
