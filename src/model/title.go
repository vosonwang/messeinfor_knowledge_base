package model

import "log"

type Title struct {
	Base
	Title string `json:"title"`
	Lang  int    `json:"lang"`
}

type Titles []Title

//根据语言查询和标题接近，并且未被占用的别名
func FindTitles(title Title) *Titles {
	var titles Titles
	if rows, err := db.Raw("SELECT d1.id, d1.title FROM doc d1 WHERE d1.deleted_at IS NULL AND d1.lang = ? AND d1.title LIKE ? AND d1.alias_id IS NULL", title.Lang, "%"+title.Title+"%").Rows(); err != nil {
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
