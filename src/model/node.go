package model

import (
	"log"
)

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

func SwapNode(down string, up string) bool {

	var d, u Doc

	err := db.Where("id = ?", down).Find(&d).Error

	err = db.Where("id = ?", up).Find(&u).Error

	if err != nil {
		return false
	}

	//数据库中number字段必须不能为unique
	d.Number, u.Number = u.Number, d.Number

	tx := db.Begin()

	err = db.Save(&d).Error

	err = db.Save(&u).Error

	if err != nil {
		tx.Rollback()
		log.Print(err)
		return false
	}

	tx.Commit()
	return true
}
