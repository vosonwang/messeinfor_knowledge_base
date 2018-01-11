package models

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"fmt"
	"messeinfor.com/messeinfor_knowledge_base/src/conf"
	"github.com/jinzhu/gorm"
	"log"
)

var db *gorm.DB

func init() {

	pg := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", conf.Host, conf.PgPort, conf.User, conf.Password, conf.Dbname)
	var err error

	if db, err = gorm.Open("postgres", pg); err != nil {
		log.Print(err)
	}

	db.LogMode(true)
}
