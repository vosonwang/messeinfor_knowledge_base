package models

import (
	"time"
	"io"
	"encoding/json"
	"github.com/satori/go.uuid"
)

type User struct {
	Id        uuid.UUID  `json:"id"`
	Username  string     `json:"username"`
	Password  string     `json:"password"`
	CreatedAt time.Time  `json:"created"`
	UpdatedAt time.Time  `json:"updated"`
	DeletedAt *time.Time `json:"deleted"`
}

func FindUser(user User) (User, bool) {

	a := db.Where(user).Find(&user)


	if a.RowsAffected == 1 {
		return user, true
	}

	return User{}, false
}

func ParseUser(body io.Reader) User {
	var user User

	if err := json.NewDecoder(body).Decode(&user); err != nil {
		panic(err)
	}

	return user
}
