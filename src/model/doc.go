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

func FindDocByName(name string, lang int) (*Doc) {
	var (
		alias Alias
		doc   Doc
	)
	//获取别名name和lang，根据name在alias表中查找alias_id
	if err := db.First(&alias, "name=?", name).Error; err != nil {
		log.Print(err)
		return nil
	}

	//再根据alias_id和lang查出doc的id
	d := db.First(&doc, "alias_id=? AND  lang =? ", alias.Id, lang)
	if d.RowsAffected == 0 {
		return &doc
	}

	if err := d.Error; err != nil {
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
