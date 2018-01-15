package model

import "log"

/*别名表*/
type Alias struct {
	Base
	Name    string `json:"name"`
	NodeKey int    `json:"nodeKey" sql:"-" `
	Docs    []Doc
}

func FindAlias(id string) *Alias {
	var alias Alias
	if err := db.First(&alias, "id=?", id).Error; err != nil {
		log.Print(err)
		return nil
	}
	return &alias
}

