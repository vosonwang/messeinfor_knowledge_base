package model

import (
	"github.com/satori/go.uuid"
	"log"
)

type Node struct {
	//一定要是Base，而不能替换成ID，不然就没法利用deleted_at
	Base
	AliasID  uuid.UUID `json:"alias_id"`
	Number   int       `json:"number" gorm:"AUTO_INCREMENT;default:0"`
	Lang     int       `json:"lang"`
	Title    string    `json:"title"`
	ParentId uuid.UUID `json:"parent_id"`
}

type Nodes []Node

func (Node) TableName() string {
	return "doc"
}

/*获取所有节点*/
func FindAllNodes(lang int) (*Nodes) {
	var nodes Nodes
	if err := db.Where("lang= ?", lang).Order("number").Find(&nodes).Error; err != nil {
		log.Print(err)
		return nil
	}
	return &nodes
}

//func SwapNode(down Alias, up Alias) bool {
//
//	tx := db.Begin()
//
//	if err := db.Exec("UPDATE alias SET node_key= ? WHERE id =?", up.NodeKey, down.Id).Error; err != nil {
//		tx.Rollback()
//		log.Print(err)
//		return false
//	}
//
//	if err := db.Exec("UPDATE alias SET node_key= ? WHERE id=?", down.NodeKey, up.Id).Error; err != nil {
//		tx.Rollback()
//		log.Print(err)
//		return false
//	}
//
//	tx.Commit()
//	return true
//}
