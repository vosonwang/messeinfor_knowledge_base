package models

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"fmt"
	"messeinfor.com/messeinfor_knowledge_base/src/conf"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func init() {

	pg := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", conf.Host, conf.PgPort, conf.User, conf.Password, conf.Dbname)
	var err error

	db, err = gorm.Open("postgres", pg)
	fmt.Print(err)

	db.LogMode(true)
}
