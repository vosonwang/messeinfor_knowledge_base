package model

import (
	"log"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

/*别名表*/
type Alias struct {
	Base
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ParentId    uuid.UUID `json:"parent_id"`
}

type Aliases []Alias

type Docs []Doc

func (alias *Alias) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.NewV4())
	return nil
}

func NewAlias(alias Alias) *Alias {
	if db.NewRecord(alias) {
		if err := db.Create(&alias).Error; err != nil {
			log.Print(err)
			return nil
		}
	} else {
		if err := db.Save(&alias).Error; err != nil {
			log.Print(err)
			return nil
		}
	}
	return &alias
}

func FindAllAlias() *Aliases {
	var aliases Aliases
	if err := db.Find(&aliases).Error; err != nil {
		log.Print(err)
		return nil
	}
	return &aliases
}

func FindAlias(id string) *Alias {
	var alias Alias
	if err := db.First(&alias, "id=?", id).Error; err != nil {
		log.Print(err)
		return nil
	}
	return &alias
}

func FindDocAlias(id string) (*DocAlias) {
	var docA DocAlias

	if err := db.Raw("SELECT * from doc where id = ?", id).Scan(&docA).Error; err != nil {
		log.Print(err)
		return nil
	}

	return &docA
}

//func UpdateDocAlias(docAlias DocAlias) (*DocAlias) {
//	var (
//		doc   Doc
//		alias Alias
//	)
//	doc = docAlias.Doc
//
//	tx := db.Begin()
//
//	if err := db.Model(&alias).Update("name", docAlias.Name).Error; err != nil {
//		tx.Rollback()
//		log.Print(err)
//		return nil
//	}
//
//	if err := db.Save(&doc).Error; err != nil {
//		tx.Rollback()
//		log.Print(err)
//		return nil
//	}
//
//	tx.Commit()
//
//	docAlias.Doc = doc
//
//	return &docAlias
//}
