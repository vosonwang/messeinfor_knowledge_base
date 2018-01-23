package model

import (
	"github.com/satori/go.uuid"
	"log"
)

type AliasTitle struct {
	Alias
	DocAliasCn uuid.UUID `json:"doc_alias_cn"`
	TitleCn    string    `json:"title_cn"`
	DocCn      uuid.UUID `json:"doc_cn"`
	DocAliasEn uuid.UUID `json:"doc_alias_en"`
	TitleEn    string    `json:"title_en"`
	DocEn      uuid.UUID `json:"doc_en"`
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

	var (
		docAliasCn DocAlias
		docAliasEn DocAlias
	)

	docAliasCn.Id = aliasTitle.DocAliasCn
	docAliasEn.Id = aliasTitle.DocAliasEn

	docAliasCn.AliasId = aliasTitle.Id
	docAliasEn.AliasId = aliasTitle.Id

	docAliasCn.DocId = aliasTitle.DocCn
	docAliasEn.DocId = aliasTitle.DocEn

	if docAliasCn.DocId != uuid.Nil {
		if db.NewRecord(docAliasCn) {
			if err := db.Create(&docAliasCn).Error; err != nil {
				log.Print(err)
				return nil
			}
		} else {
			if err := db.Save(&docAliasCn).Error; err != nil {
				log.Print(err)
				return nil
			}
		}
	}

	if docAliasEn.DocId != uuid.Nil {
		if db.NewRecord(docAliasEn) {

			if err := db.Create(&docAliasEn).Error; err != nil {
				log.Print(err)
				return nil
			}
		} else {
			if err := db.Save(&docAliasEn).Error; err != nil {
				log.Print(err)
				return nil
			}
		}
	}

	aliasTitle.DocAliasCn = docAliasCn.Id
	aliasTitle.DocAliasEn = docAliasEn.Id

	return &aliasTitle
}

func FindAllAliasTitle() *AliasTitles {
	var ats AliasTitles
	if rows, err := db.Raw("SELECT  a.id as id,a.created_at as created_at,a.updated_at as updated_at,a.deleted_at as deleted_at,a.name as name ,a.description as description,a.parent_id as parent_id, dax.id  AS doc_alias_cn,  dax.title  AS title_cn,  dax.doc_id AS doc_cn,  (SELECT da.id AS doc_alias_en FROM doc d INNER JOIN doc_alias da ON d.id = da.doc_id WHERE d.lang = 1 AND a.id = da.alias_id LIMIT 1),  (SELECT d.title AS title_en FROM doc d INNER JOIN doc_alias da ON d.id = da.doc_id WHERE d.lang = 1 AND a.id = da.alias_id LIMIT 1),  (SELECT da.doc_id AS doc_cn FROM doc d INNER JOIN doc_alias da ON d.id = da.doc_id WHERE d.lang = 1 AND a.id = da.alias_id LIMIT 1)FROM alias a LEFT JOIN (SELECT da.id  AS id, d.title  AS title, d.lang AS lang, da.doc_id AS doc_id, da.alias_id AS alias_id FROM doc d INNER JOIN doc_alias da ON d.id = da.doc_id WHERE d.lang = 0) AS dax ON a.id = dax.alias_id WHERE a.deleted_at is null").Rows(); err != nil {
		log.Print(err)
		return nil
	} else {
		for rows.Next() {
			var at AliasTitle
			rows.Scan(&at.Id, &at.CreatedAt, &at.UpdatedAt, &at.DeletedAt, &at.Name, &at.Description, &at.ParentId, &at.DocAliasCn, &at.TitleCn, &at.DocCn, &at.DocAliasEn, &at.TitleEn, &at.DocEn)
			ats = append(ats, at)
		}
	}
	return &ats
}

func DeleteAliasTitle(id uuid.UUID) bool {
	var (
		docAlias DocAlias
		alias    Alias
	)
	alias.Id = id
	docAlias.AliasId = id

	tx := db.Begin()

	if err := tx.Delete(&alias).Error; err != nil {
		tx.Rollback()
		log.Print(err)
		return false
	}

	if err := tx.Delete(&docAlias).Error; err != nil {
		tx.Rollback()
		log.Print(err)
		return false
	}

	tx.Commit()
	return true
}

func FindAliasTitle(id string) *AliasTitle {
	var aliasTitle AliasTitle
	if err := db.Raw("SELECT  a.id as id,a.created_at as created_at,a.updated_at as updated_at,a.deleted_at as deleted_at,a.name as name ,a.description as description,a.parent_id as parent_id, dax.id  AS doc_alias_cn,  dax.title  AS title_cn,  dax.doc_id AS doc_cn, (SELECT da.id AS doc_alias_en FROM doc d INNER JOIN doc_alias da ON d.id = da.doc_id WHERE d.lang = 1 AND a.id = da.alias_id LIMIT 1),  (SELECT d.title AS title_en FROM doc d INNER JOIN doc_alias da ON d.id = da.doc_id WHERE d.lang = 1 AND a.id = da.alias_id LIMIT 1),  (SELECT da.doc_id AS doc_cn FROM doc d INNER JOIN doc_alias da ON d.id = da.doc_id WHERE d.lang = 1 AND a.id = da.alias_id LIMIT 1)FROM alias a LEFT JOIN (SELECT da.id  AS id, d.title  AS title, d.lang AS lang, da.doc_id AS doc_id, da.alias_id AS alias_id FROM doc d INNER JOIN doc_alias da ON d.id = da.doc_id WHERE d.lang = 0) AS dax ON a.id = dax.alias_id WHERE a.deleted_at is null and a.id=?", id).Scan(&aliasTitle).Error; err != nil {
		log.Print(err)
		return nil
	}
	return &aliasTitle
}
