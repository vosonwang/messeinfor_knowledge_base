package model

import (
	"github.com/satori/go.uuid"
	"log"
)

type Doc struct {
	Base
	Lang     int       `json:"lang"`
	Text     string    `json:"text"`
	Title    string    `json:"title"`
	ParentId uuid.UUID `json:"parent_id"`
	AliasId  uuid.UUID `json:"alias_id"`
}


func FindDoc(id string) (*Doc) {
	var doc Doc
	if err := db.First(&doc, "id=?", id).Error; err != nil {
		log.Print(err)
		return nil
	}
	return &doc
}

func DeleteDoc(doc Doc) bool {
	var alias Alias
	alias.Id = doc.AliasId

	tx := db.Begin()

	if err := tx.Delete(&doc).Error; err != nil {
		tx.Rollback()
		log.Print(err)
		return false
	}

	if err := tx.Delete(&alias).Error; err != nil {
		tx.Rollback()
		log.Print(err)
		return false
	}

	tx.Commit()

	return true
}
