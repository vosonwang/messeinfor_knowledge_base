package model


/*别名表*/
type Alias struct {
	Base
	Name    string `json:"name"`
	NodeKey int    `json:"nodeKey" sql:"-" `
	Docs    []Doc
}

func FindAlias(id string) (Alias, error) {
	var alias Alias
	err := db.First(&alias, "id=?", id).Error
	return alias, err
}

//func FindDocAlias(id string) (DocAlias, error) {
//	var docA DocAlias
//	if err := db.Raw("SELECT d.*,a.node_key,a.name FROM doc d INNER JOIN alias a ON d.alias_id = a.id WHERE d.deleted_at IS NULL AND d.id = ?'", id).Scan(&docA).Error; err != nil {
//		return docA, err
//	}
//	return docA, nil
//}