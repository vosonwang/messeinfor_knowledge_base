package model

import (
	"github.com/satori/go.uuid"
	"log"
	"github.com/jinzhu/gorm"
)


func (doc *Doc) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.NewV4())
	return nil
}

func NewDoc(doc *Doc) bool {
	if db.NewRecord(doc) {
		if err := db.Create(&doc).Error; err != nil {
			log.Print(err)
			return false
		}

	} else {
		if err := db.Save(&doc).Error; err != nil {
			log.Print(err)
			return false
		}

	}
	return true
}

func FindDoc(id string) (*Doc) {
	var doc Doc
	if err := db.First(&doc, "id=?", id).Error; err != nil {
		log.Print(err)
		return nil
	}
	return &doc
}

func FindAllDoc() *Docs {
	var docs Docs
	err := db.Find(&docs).Error
	if err != nil {
		log.Print(err)
		return nil
	}
	return &docs

}

func FindDocByAlias(aliasId uuid.UUID, lang int) *Doc {
	var doc Doc

	if err := db.First(&doc, "alias_id=? AND  lang =? ", aliasId, lang).Error; err != nil {
		log.Print(err)
		return nil
	}
	return &doc
}

func DeleteDoc(doc Doc) bool {
	tx := db.Begin()
	if err := tx.Delete(&doc).Error; err != nil {
		tx.Rollback()
		log.Print(err)
		return false
	}

	if err := tx.Exec("update doc set alias_id = null where alias_id = ?", doc.AliasID).Error; err != nil {
		tx.Rollback()
		log.Print(err)
		return false
	}

	tx.Commit()

	return true
}
