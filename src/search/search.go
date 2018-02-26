package search

import (
	"github.com/olivere/elastic"
	"fmt"
	"messeinfor.com/messeinfor_knowledge_base/src/model"
	"net/http"
	"messeinfor.com/messeinfor_knowledge_base/src/conf"
	"messeinfor.com/messeinfor_knowledge_base/src/handler"
	"github.com/gorilla/mux"
)

func MultiMatch(w http.ResponseWriter, r *http.Request) {
	//过滤特殊字符
	vars := mux.Vars(r)

	q := elastic.NewMultiMatchQuery(vars["words"], "title", "text").Fuzziness("AUTO")
	searchResult, err := es.Search().
		Index(conf.E.Index).
		Query(q).
		From(0).Size(100).
		Highlight(
		elastic.NewHighlight().
			Fields(elastic.NewHighlighterField("text").
			PreTags("<b>").PostTags("</b>"),
			elastic.NewHighlighterField("title").
				PreTags("<b>").PostTags("</b>"), ).NumOfFragments(1)).
		Pretty(true).
		Do(ctx)

	if err != nil {
		panic(err)
	}

	handler.JsonResponse(w, searchResult)

}

func AddDoc(doc *model.Doc) {
	d, err := es.Index().
		Index(conf.E.Index).
		Type(conf.E.Type).
		Id(doc.Id.String()).
		BodyJson(doc).
		Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}

	fmt.Printf("Indexed doc %s to index %s, type %s\n", d.Id, d.Index, d.Type)
}

func UpdateDoc() {
	//	Script(elastic.NewScriptInline("ctx._source.retweets += params.num").Lang("painless").Param("num", 1)).
	//	Upsert(map[string]interface{}{"retweets": 0}).
	//	Do(ctx)
	//if err != nil {
	//	// Handle error
	//	panic(err)
	//}
	//fmt.Printf("New version of tweet %q is now %d\n", update.Id, update.Version)
}
