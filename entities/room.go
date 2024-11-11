package entities

import (
	"fmt"
	"github.com/google/uuid"
)

type Room struct {
	ID         string
	Attendants map[string]*Attendant
}

func NewRoom() *Room {
	return &Room{
		ID:         uuid.New().String(),
		Attendants: make(map[string]*Attendant),
	}
}

func (r *Room) AddAttendant(a *Attendant) {
	fmt.Println("adding attendant: ", a)
}
