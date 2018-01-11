package models

import (
	"time"
	"github.com/satori/go.uuid"
	"fmt"
	"log"
)

type Node struct {
	Id        uuid.UUID  `json:"id"`
	Title     string     `json:"title"`
	NodeKey   int        `json:"nodeKey" gorm:"auto_increment; primary_key; unique"`
	Lang      int        `json:"lang"`
	ParentId  uuid.UUID  `json:"parent_id"`
	CreatedAt time.Time  `json:"created"`
	UpdatedAt time.Time  `json:"updated"`
	DeletedAt *time.Time `json:"deleted"`
}

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

type Nodes []Node

func (Node) TableName() string {
	return "docs"
}

func FindNodes(lang int) (Nodes) {
	var nodes Nodes
	/*按照node_key排序，以便前端按照此顺序由上到下排列*/
	if err := db.Where("lang = ?", lang).Order("node_key").Find(&nodes).Error; err == nil {
		return nodes
	} else {
		fmt.Print(err)
		return nil
	}
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

func Swap(a Doc, b Doc) error {
	a.NodeKey, b.NodeKey = b.NodeKey, a.NodeKey

	e := db.Save(&a).Error
	f := db.Save(&b).Error
	if e != nil {
		return e
	}
	if f != nil {
		return f
	}
	return nil

}
