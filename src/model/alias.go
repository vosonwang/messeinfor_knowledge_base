package model

import (
	"log"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

func (alias *Alias) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.NewV4())
	return nil
}

func FindAlias(id string) (*Alias) {
	var a Alias
	if err := db.First(&a, "id=?", id).Error; err != nil {
		log.Print(err)
		return nil
	}
	return &a
}

func FindAliasByName(name string) (*Alias) {
	var a Alias
	if err := db.First(&a, "name=?", name).Error; err != nil {
		log.Print(err)
		return nil
	}
	return &a
}

/*查询和描述相关的，所有未使用的别名列表*/
func FindAliasByDesc(description string) *Aliases {

	var aliases Aliases
	if rows, err := db.Raw("select a.* from alias a LEFT JOIN doc d on a.id=d.alias_id WHERE a.deleted_at IS NULL AND d.id is NULL AND a.description LIKE ?", "%"+description+"%").Rows(); err != nil {
		log.Print(err)
		return nil
	} else {
		for rows.Next() {
			var a Alias
			db.ScanRows(rows, &a)
			aliases = append(aliases, a)
		}
	}
	return &aliases
}
