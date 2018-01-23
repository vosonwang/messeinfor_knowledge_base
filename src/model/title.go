package model

import "log"

type Title struct {
	Base
	Title string `json:"title"`
	Lang  int    `json:"lang"`
}

type Titles []Title

func (Title) TableName() string {
	return "doc"
}

func FindTitles(title Title) *Titles {
	var titles Titles
	if rows, err := db.Raw("select id,title from doc where deleted_at is null AND lang = ? AND title LIKE  ? AND id not in (select doc_id from doc_alias where deleted_at is null)", title.Lang, "%"+title.Title+"%").Rows(); err != nil {
		log.Print(err)
		return nil
	} else {
		for rows.Next() {
			var t Title
			rows.Scan(&t.Id, &t.Title)
			titles = append(titles, t)
		}
	}
	return &titles
}
