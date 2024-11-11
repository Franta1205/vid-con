package entities

import "github.com/google/uuid"

type Attendant struct {
	ID   string
	Name string
}

func NewAttendant(name string) *Attendant {
	return &Attendant{
		ID:   uuid.New().String(),
		Name: name,
	}
}
