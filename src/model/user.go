package model

import "log"

type User struct {
	Base
	Username string `json:"username"`
	Password string `json:"password"`
}

func FindUser(user User) (*User) {

	a := db.Where(user).Find(&user)

	if a.Error == nil && a.RowsAffected == 1 {
		return &user
	}
	log.Print(a.Error)
	return nil
}
