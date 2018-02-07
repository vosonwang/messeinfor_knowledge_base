package cache

import (
	"log"
	"messeinfor.com/messeinfor_knowledge_base/src/model"
	"github.com/RedisLabs/redisearch-go/redisearch"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"messeinfor.com/messeinfor_knowledge_base/src/handler"
)

//将文档存入rSearch
func AddDoc(doc model.Doc) bool {
	d := redisearch.NewDocument((doc.Id).String(), 1.0)

	d.Set("title", doc.Title).
		Set("body", doc.Text).
		Set("alias_id", doc.AliasID).
		Set("updated", doc.UpdatedAt)

	if doc.Lang == 0 {
		d.Set("language", "chinese")
	}

	// Index the document. The API accepts multiple documents at a time
	if err := rSearch.IndexOptions(redisearch.DefaultIndexingOptions, d); err != nil {
		log.Print(err)
		return false
	}

	return true
}

func SearchDoc(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	text := vars["text"]
	lang := vars["lang"]

	q := redisearch.NewQuery(text).
		Highlight([]string{"body"}, "<b>", "</b>").
		SummarizeOptions(redisearch.SummaryOptions{
		Fields:      []string{"body"},
		FragmentLen: 50,
		NumFragments:2})

	if lang == "0" {
		q.SetLanguage("chinese")
	}

	docs, total, err := rSearch.Search(q)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "找不到文档！")
	} else {
		fmt.Print(total)
		handler.JsonResponse(w, docs)
	}

}
