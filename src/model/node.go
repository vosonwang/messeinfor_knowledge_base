package model

import (
	"github.com/satori/go.uuid"
	"log"
)

type Node struct {
	Id       uuid.UUID `json:"id"`
	Title    string    `json:"title"`
	ParentId uuid.UUID `json:"parent_id"`
	NodeKey  int       `json:"nodeKey"`
	AliasId  uuid.UUID `json:"alias_id"`
}

type Nodes []Node

/*获取所有节点*/
func FindNodes(lang int) (Nodes, bool) {
	var (
		nodes Nodes
		node  Node
	)
	if rows, err := db.Raw("SELECT d.id,d.title,d.parent_id,a.node_key,a.id FROM doc d INNER JOIN alias a ON d.alias_id=a.id WHERE d.deleted_at IS NULL AND  d.lang = ? ORDER BY a.node_key", lang).Rows(); err != nil {
		return nil, false
	} else {
		defer rows.Close()
		for rows.Next() {
			rows.Scan(&node.Id, &node.Title, &node.ParentId, &node.NodeKey, &node.AliasId)
			nodes = append(nodes, node)
		}
	}
	return nodes, true
}

func SwapNode(down Alias, up Alias) bool {
	down.NodeKey, up.NodeKey = up.NodeKey, down.NodeKey
	tx := db.Begin()

	if err := db.Save(&down).Error; err != nil {
		log.Print(err)
		return false
	}

	if err := db.Save(&down).Error; err != nil {
		log.Print(err)
		return false
	}

	tx.Commit()
	return false
}
