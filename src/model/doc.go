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

type DocAlias struct {
	Doc
	NodeKey int    `json:"nodeKey"`
	Name    string `json:"name"`
}

type Docs []Doc

func AddDoc(doc Doc) (Doc, error) {
	var alias Alias
	alias.Id = uuid.NewV4()
	doc.Id = uuid.NewV4()
	doc.AliasId = alias.Id

	tx := db.Begin()

	if err := tx.Create(&alias).Error; err != nil {
		tx.Rollback()
		log.Print(err)
		return Doc{}, err
	}

	if err := tx.Create(&doc).Error; err != nil {
		tx.Rollback()
		log.Print(err)
		return Doc{}, err
	}

	tx.Commit()

	return doc, nil
}

func FindDoc(id string) (Doc, error) {
	var doc Doc
	err := db.First(&doc, "id=?", id).Error
	return doc, err
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

func UpdateDoc(doc Doc) (Doc, error) {
	if err := db.Save(&doc).Error; err != nil {
		return doc, err
	}
	//可能是postgres的驱动语言即lib/pq，更新插入时，updated_at created_at 都是cst时间，这导致我不得不在此处手动替换updated_at为utc时间

	return doc, nil
}
