package models

import (
	"time"
	"github.com/satori/go.uuid"
	"fmt"
	"log"
)

/*TODO 由于是中英文双语，所以有id和lang 这种联名主键，因此更新时需要考虑语言*/
type Doc struct {
	Id        uuid.UUID  `json:"id"`
	Title     string     `json:"title"`
	Text      string     `json:"text" gorm:"column:doc"`
	NodeKey   int        `json:"nodeKey" gorm:"auto_increment; primary_key; unique"`
	Lang      int        `json:"lang"`
	ParentId  uuid.UUID  `json:"parent_id"`
	CreatedAt time.Time  `json:"created"`
	UpdatedAt time.Time  `json:"updated"`
	DeletedAt *time.Time `json:"deleted"`
}



func AddDoc(doc Doc) (uuid.UUID) {
	doc.Id = uuid.NewV4()
	if err := db.Create(&doc).Error; err != nil {
		fmt.Print(err)
		return uuid.Nil
	}
	return doc.Id
}

func FindDoc(id string) (Doc, error) {
	var doc Doc
	err := db.First(&doc, "id=?", id).Error
	return doc, err
}

func DeleteDoc(doc Doc) bool {
	if err := db.Delete(&doc).Error; err != nil {
		log.Print(err)
		return false
	}
	return true
}

func UpdateDoc(doc Doc) (time.Time, error) {
	cst := doc.UpdatedAt //使用存进数据库之前的时间，作为返回前台的修改时间
	if err := db.Save(&doc).Error; err != nil {
		return cst, err
	}
	return cst, nil
}


