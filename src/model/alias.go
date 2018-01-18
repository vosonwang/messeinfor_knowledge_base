package model

import (
	"log"
	"github.com/satori/go.uuid"
)

/*别名表*/
type Alias struct {
	Base
	Name    string `json:"name"`
	NodeKey int    `json:"nodeKey" sql:"-" `
	Docs    []Doc
}

type DocAlias struct {
	Doc
	Name string `json:"name"`
}

type Docs []Doc

func FindAlias(id string) *Alias {
	var alias Alias
	if err := db.First(&alias, "id=?", id).Error; err != nil {
		log.Print(err)
		return nil
	}
	return &alias
}

func FindDocAlias(id uuid.UUID) (*DocAlias) {
	var docA DocAlias

	if err := db.Raw("SELECT d.*,a.name FROM doc d INNER JOIN alias a ON d.alias_id = a.id WHERE d.deleted_at IS NULL AND d.id = ?", id).Scan(&docA).Error; err != nil {
		log.Print(err)
		return nil
	}

	return &docA
}

func UpdateDocAlias(docAlias DocAlias) (*DocAlias) {
	var (
		doc   Doc
		alias Alias
	)
	doc = docAlias.Doc
	alias.Id = docAlias.AliasId

	tx := db.Begin()

	if err := db.Model(&alias).Update("name", docAlias.Name).Error; err != nil {
		tx.Rollback()
		log.Print(err)
		return nil
	}

	if err := db.Save(&doc).Error; err != nil {
		tx.Rollback()
		log.Print(err)
		return nil
	}

	tx.Commit()

	docAlias.Doc = doc

	return &docAlias
}
