package server

import (
	"fmt"
	"log"
	"net/http"
)

type User struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Experience int64  `json:"experience"`
}

func NewUser(name, surname string, id, experience int64) *User {
	return &User{
		Name:       name,
		Surname:    surname,
		ID:         id,
		Experience: experience,
	}
}

func (u *User) saveToDB() {
	log.Printf("Succesufelly save to DB %v", u.ID)
}

func sendBadNotification(text string, w http.ResponseWriter) {
	w.Write([]byte(fmt.Sprintf("You operation was unsucceed: %v", text)))
}

func sendGoodNotification(w http.ResponseWriter) error {
	_, err := w.Write([]byte("You operation was succeed!"))
	if err != nil {
		return err
	}

	return nil
}
