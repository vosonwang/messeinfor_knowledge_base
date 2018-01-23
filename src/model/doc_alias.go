package model

import (
	"github.com/satori/go.uuid"
	"github.com/jinzhu/gorm"
)

type DocAlias struct {
	Base
	AliasId uuid.UUID `json:"alias_id"`
	DocId   uuid.UUID `json:"doc_id"`
}

func (docAlias *DocAlias) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.NewV4())
	return nil
}
