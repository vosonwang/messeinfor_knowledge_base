package model

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"fmt"
	"messeinfor.com/messeinfor_knowledge_base/src/conf"
	"github.com/jinzhu/gorm"
	"log"
	"github.com/satori/go.uuid"
	"time"
)

type Base struct {
	Id        uuid.UUID  `json:"id"`
	CreatedAt time.Time  `json:"created"`
	UpdatedAt time.Time  `json:"updated"`
	DeletedAt *time.Time `json:"deleted"`
}



var db *gorm.DB

func init() {

	pg := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", conf.PgHost, conf.PgPort, conf.User, conf.Password, conf.Dbname)
	var err error

	if db, err = gorm.Open("postgres", pg); err != nil {
		log.Print(err)
	}

	//set this to true, `User`'s default table name will be `user`, table name setted with `TableName` won't be affected
	db.SingularTable(true)

	db.LogMode(true)
}
