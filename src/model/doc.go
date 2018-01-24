package model

import (
	"github.com/satori/go.uuid"
	"log"
	"github.com/jinzhu/gorm"
)

type Doc struct {
	Base
	Number   int       `json:"number" gorm:"AUTO_INCREMENT;default:0"`
	Lang     int       `json:"lang"`
	Text     string    `json:"text"`
	Title    string    `json:"title"`
	AliasID  uuid.UUID `json:"alias_id"`
	ParentId uuid.UUID `json:"parent_id"`
	Creator  uuid.UUID `json:"creator"`
	Updater  uuid.UUID `json:"updater"`
}

func (doc *Doc) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.NewV4())
	return nil
}

func NewDoc(doc Doc) *Doc {
	if db.NewRecord(doc) {
		if err := db.Create(&doc).Error; err != nil {
			log.Print(err)
			return nil
		}
	} else {
		if err := db.Save(&doc).Error; err != nil {
			log.Print(err)
			return nil
		}
	}
	return &doc
}

func FindDoc(id string) (*Doc) {
	var doc Doc
	if err := db.First(&doc, "id=?", id).Error; err != nil {
		log.Print(err)
		return nil
	}
	return &doc
}

func FindDocByAlias(aliasId string, lang int) *Doc {
	var doc Doc

	if err := db.First(&doc, "alias_id=? AND  lang =? ", aliasId, lang).Error; err != nil {
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

	if err := db.First(&doc, "alias_id=? AND  lang =? ", alias.Id, lang).Error; err != nil {
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
