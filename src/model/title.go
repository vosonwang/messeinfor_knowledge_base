package model

import (
	"log"
	"github.com/satori/go.uuid"
)

type Title struct {
	Base
	Title string `json:"title"`
	Lang  int    `json:"lang"`
}

type Titles []Title

type AliasTitle struct {
	Alias
	DocCn   uuid.UUID `json:"doc_cn" `
	TitleCn string    `json:"title_cn" `
	DocEn   uuid.UUID `json:"doc_en" `
	TitleEn string    `json:"title_en" `
}

type AliasTitles []AliasTitle

func NewAliasTitle(aliasTitle AliasTitle) *AliasTitle {

	if db.NewRecord(aliasTitle) {
		if err := db.Create(&aliasTitle.Alias).Error; err != nil {
			log.Print(err)
			return nil
		}
	} else {
		if err := db.Save(&aliasTitle.Alias).Error; err != nil {
			log.Print(err)
			return nil
		}
	}

	/*不管AliasTitle中存不存在中英文文档信息，都清空，稍后如果信息存在，则重新添加（不管alias_id有没有变更）*/
	if err := db.Exec("update doc set alias_id = null where alias_id = ?", aliasTitle.Id).Error; err != nil {
		log.Print(err)
		return nil
	}

	if aliasTitle.DocCn != uuid.Nil {
		var docCn Doc
		docCn.Id = aliasTitle.DocCn

		if err := db.Model(&docCn).Update("alias_id", aliasTitle.Id).Error; err != nil {
			log.Print(err)
			return nil
		}
	}

	if aliasTitle.DocEn != uuid.Nil {
		var docEn Doc

		docEn.Id = aliasTitle.DocEn

		if err := db.Model(&docEn).Update("alias_id", aliasTitle.Id).Error; err != nil {
			log.Print(err)
			return nil
		}
	}

	return &aliasTitle
}

func FindAllAliasTitle() *AliasTitles {
	var ats AliasTitles
	if rows, err := db.Raw("SELECT a.*, ( SELECT d.id AS doc_cn FROM doc d WHERE d.alias_id = a.id AND d.lang = 0 LIMIT 1 ) , ( SELECT d1.title AS title_cn FROM doc d1 WHERE d1.alias_id = a.id AND d1.lang = 0 LIMIT 1 ) , ( SELECT d2.id AS doc_en FROM doc d2 WHERE d2.alias_id = a.id AND d2.lang = 1 LIMIT 1 ) , ( SELECT d3.title AS title_en FROM doc d3 WHERE d3.alias_id = a.id AND d3.lang = 1 LIMIT 1 ) FROM alias a WHERE a.deleted_at IS NULL ORDER BY number").Rows(); err != nil {
		log.Print(err)
		return nil
	} else {
		for rows.Next() {
			var at AliasTitle
			db.ScanRows(rows, &at)
			ats = append(ats, at)
		}
	}
	return &ats
}

func DeleteAliasTitle(id uuid.UUID) bool {
	var alias Alias

	alias.Id = id

	tx := db.Begin()

	if err := tx.Delete(&alias).Error; err != nil {
		tx.Rollback()
		log.Print(err)
		return false
	}

	if err := tx.Exec("update doc set alias_id = null where alias_id = ?", id).Error; err != nil {
		tx.Rollback()
		log.Print(err)
		return false
	}

	tx.Commit()
	return true
}

func FindAliasTitle(id string) *AliasTitle {
	var aliasTitle AliasTitle
	if err := db.Raw("SELECT a.* , ( SELECT d.id AS doc_cn FROM doc d WHERE d.alias_id = a.id AND d.lang = 0 LIMIT 1 ) , ( SELECT d1.title AS title_cn FROM doc d1 WHERE d1.alias_id = a.id AND d1.lang = 0 LIMIT 1 ) , ( SELECT d2.id AS doc_en FROM doc d2 WHERE d2.alias_id = a.id AND d2.lang = 1 LIMIT 1 ) , ( SELECT d3.title AS title_en FROM doc d3 WHERE d3.alias_id = a.id AND d3.lang = 1 LIMIT 1 ) FROM alias a WHERE a.deleted_at IS NULL AND a.id=?", id).Scan(&aliasTitle).Error; err != nil {
		log.Print(err)
		return nil
	}
	return &aliasTitle
}

//根据语言查询和标题接近，并且未被占用的别名
func FindTitles(value string,lang int) *Titles {
	var titles Titles
	if rows, err := db.Raw("SELECT d1.id, d1.title FROM doc d1 WHERE d1.deleted_at IS NULL AND d1.lang = ? AND d1.title LIKE ? AND (d1.alias_id IS NULL OR d1.alias_id = '00000000-0000-0000-0000-000000000000')", lang, "%"+value+"%").Rows(); err != nil {
		log.Print(err)
		return nil
	} else {
		for rows.Next() {
			var t Title
			rows.Scan(&t.Id, &t.Title)
			titles = append(titles, t)
		}
	}
	return &titles
}
