package models

import (
	"time"
	"github.com/satori/go.uuid"
	"fmt"
)

type Node struct {
	Id        uuid.UUID  `json:"id"  gorm:"primary_key"`
	Title     string     `json:"title"`
	NodeKey   int        `json:"nodeKey"`
	Lang      int        `json:"lang"  gorm:"primary_key"`
	ParentId  uuid.UUID  `json:"parent_id"`
	CreatedAt time.Time  `json:"created"`
	UpdatedAt time.Time  `json:"updated"`
	DeletedAt *time.Time `json:"deleted"`
}

type Nodes []Node

func (Node) TableName() string {
	return "docs"
}

func FindNodes(lang int) (Nodes) {
	var nodes Nodes
	/*按照node_key排序，以便前端按照此顺序由上到下排列*/
	if err := db.Where("lang = ?", lang).Order("node_key").Find(&nodes).Error; err == nil {
		return nodes
	} else {
		fmt.Print(err)
		return nil
	}
}

/*TODO
如果a,b 中英文都存在，则都交换

如果a,b中文都存在 a或b中有一个英文不存在，则中文交换  英文不动

如果a,b英文都存在，同上
*/
func FindNode(id string,lang int) Node{
	var node Node

	return node
}

func Swap(a Node, b Node) error {
	c := a.NodeKey
	c, b.NodeKey = b.NodeKey, c

	if d := db.Model(&a).Update("node_key", 999999).Error; d != nil {
		return d
	}

	if f := db.Save(&b).Error; f != nil {
		return f
	}

	if e := db.Model(&a).Update("node_key", c).Error; e != nil {
		return e
	}

	return nil

}