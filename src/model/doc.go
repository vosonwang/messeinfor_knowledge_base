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

func GetAllDocId() *[]string {
	var Ids []string

	if rows, err := db.Raw("SELECT id FROM doc where deleted_at is null").Rows(); err != nil {
		log.Print(err)
		return nil
	} else {
		for rows.Next() {
			var id string
			rows.Scan(&id)
			Ids = append(Ids, id)
		}
		return &Ids
	}

}

func FindDoc(id string) (*Doc) {
	var doc Doc
	if err := db.First(&doc, "id=?", id).Error; err != nil {
		log.Print(err)
		return nil
	}
	return &doc
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
