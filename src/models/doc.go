package models

import (
	"time"
	"github.com/satori/go.uuid"
	"io"
	"encoding/json"
	"fmt"
	"log"
)

type Doc struct {
	Id        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	NodeKey   int       `json:"nodeKey" gorm:"AUTO_INCREMENT"`
	Lang      int       `json:"lang"`
	ParentId  uuid.UUID `json:"parent_id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type Docs []Doc

func GetDocs(lang int) (Docs) {
	var docs Docs
	/*按照node_key排序，以便前端按照此顺序由上到下排列*/
	if err := db.Where("lang = ?", lang).Order("node_key").Find(&docs).Error; err == nil {
		return docs
	} else {
		fmt.Print(err)
		return nil
	}
}

func ParseNode(body io.Reader) (map[string]interface{}) {
	var a interface{}
	if err := json.NewDecoder(body).Decode(&a); err != nil {
		log.Print(err)
		return nil
	}
	return a.(map[string]interface{})
}

func AddDoc(node interface{}) (uuid.UUID) {
	a := node.(map[string]interface{})
	var b Doc
	b.Id = uuid.NewV4()
	b.Title = a["title"].(string)
	b.ParentId, _ = uuid.FromString(a["parent_id"].(string))
	b.Lang = int(a["lang"].(float64))
	if err := db.Create(&b).Error; err != nil {
		fmt.Print(err)
		return uuid.Nil
	}
	return b.Id
}
