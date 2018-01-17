package model

import (
	"github.com/satori/go.uuid"
	"log"
)

type Node struct {
	Base
	Lang     int       `json:"lang"`
	Title    string    `json:"title"`
	ParentId uuid.UUID `json:"parent_id"`
	AliasId  uuid.UUID `json:"alias_id"`
}

type Nodes []Node

func (Node) TableName() string {
	return "doc"
}

func AddNode(node Node) (*Node) {
	var alias Alias
	alias.Id = uuid.NewV4()
	node.Id = uuid.NewV4()
	node.AliasId = alias.Id

	tx := db.Begin()

	if err := tx.Create(&alias).Error; err != nil {
		tx.Rollback()
		log.Print(err)
		return nil
	}

	if err := tx.Create(&node).Error; err != nil {
		tx.Rollback()
		log.Print(err)
		return nil
	}

	tx.Commit()

	return &node
}

/*获取所有节点*/
func FindNodes(lang int) (*Nodes, bool) {
	var (
		nodes Nodes
		node  Node
	)
	if rows, err := db.Raw("SELECT d.id,d.title,d.parent_id,d.lang,a.id FROM doc d INNER JOIN alias a ON d.alias_id=a.id WHERE d.deleted_at IS NULL AND  d.lang = ? ORDER BY a.node_key", lang).Rows(); err != nil {
		return nil, false
	} else {
		defer rows.Close()
		for rows.Next() {
			rows.Scan(&node.Id, &node.Title, &node.ParentId, &node.Lang, &node.AliasId)
			nodes = append(nodes, node)
		}
	}
	return &nodes, true
}

func SwapNode(down Alias, up Alias) bool {

	tx := db.Begin()

	if err := db.Exec("UPDATE alias SET node_key= ? WHERE id =?", up.NodeKey, down.Id).Error; err != nil {
		tx.Rollback()
		log.Print(err)
		return false
	}

	if err := db.Exec("UPDATE alias SET node_key= ? WHERE id=?", down.NodeKey, up.Id).Error; err != nil {
		tx.Rollback()
		log.Print(err)
		return false
	}

	tx.Commit()
	return true
}
