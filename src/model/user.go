package model

import "log"

type User struct {
	Base
	Username string `json:"username"`
	Password string `json:"password"`
	Docs     []Doc
}

func FindUser(user User) (*User) {

	a := db.Where(user).Find(&user)

	if a.Error == nil && a.RowsAffected == 1 {
		return &user
	}
	/*TODO 此处无法明确判断是用户不存在 还是 连接数据库失败*/
	log.Print(a.Error)
	return nil
}
