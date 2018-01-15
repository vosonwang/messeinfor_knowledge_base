package model

import "log"

type User struct {
	Base
	Username string `json:"username"`
	Password string `json:"password"`
}

func FindUser(user User) (User, bool) {

	a := db.Where(user).Find(&user)

	if a.Error == nil && a.RowsAffected == 1 {
		return user, true
	}
	log.Print(a.Error)
	return User{}, false
}
