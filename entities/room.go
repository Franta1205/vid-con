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
	r.Attendants[a.ID] = a
}

func (r *Room) RemoveAttendant(a *Attendant) {
	fmt.Println("removing attendant", a)
	if _, exist := r.Attendants[a.ID]; exist {
		delete(r.Attendants, a.ID)
		fmt.Println("attendant removed", a.Name)
	} else {
		fmt.Println("attendant not found")
	}
}
