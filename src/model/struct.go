package model

import (
	"github.com/satori/go.uuid"
)

type Base struct {
	Id        uuid.UUID `json:"id"`
	CreatedAt int64     `json:"created"`
	UpdatedAt int64     `json:"updated"`
	DeletedAt int64     `json:"deleted,omitempty"`
}

type Doc struct {
	Base
	Number   int       `json:"number" gorm:"AUTO_INCREMENT;default:0"`
	Lang     int       `json:"lang"`
	Text     string    `json:"text"`
	Title    string    `json:"title"`
	AliasID  uuid.UUID `json:"alias_id,omitempty"`
	ParentId uuid.UUID `json:"parent_id"`
	Creator  uuid.UUID `json:"creator"`
	Updater  uuid.UUID `json:"updater"`
}

/*别名表*/
type Alias struct {
	Base
	Number      int       `json:"number" gorm:"AUTO_INCREMENT;default:0"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ParentId    uuid.UUID `json:"parent_id"`
}

type Aliases []Alias

type Docs []Doc

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

type Title struct {
	Base
	Title string `json:"title"`
	Lang  int    `json:"lang"`
}

type Titles []Title

type AliasTitle struct {
	Alias
	DocCn   uuid.UUID `json:"doc_cn" `
	TitleCn string    `json:"title_cn" `
	DocEn   uuid.UUID `json:"doc_en" `
	TitleEn string    `json:"title_en" `
}

type AliasTitles []AliasTitle

type User struct {
	Base
	Username string `json:"username"`
	Password string `json:"password"`
}
