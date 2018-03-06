package model

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"fmt"
	"messeinfor.com/messeinfor_knowledge_base/src/conf"
	"github.com/jinzhu/gorm"
	"log"
)

var db *gorm.DB

func init() {

	pg := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", conf.P.Host, conf.P.Port, conf.P.User, conf.P.Password, conf.P.Db)
	var err error

	if db, err = gorm.Open("postgres", pg); err != nil {
		log.Print(err)
	}

	//set this to true, `User`'s default table name will be `user`, table name setted with `TableName` won't be affected
	db.SingularTable(true)

	db.LogMode(true)
}
