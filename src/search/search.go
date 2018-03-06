package search

import (
	"github.com/olivere/elastic"
	"fmt"
	"messeinfor.com/messeinfor_knowledge_base/src/model"
	"net/http"
	"messeinfor.com/messeinfor_knowledge_base/src/conf"
	"github.com/gorilla/mux"
	"context"
	"messeinfor.com/messeinfor_knowledge_base/src/util"
)

/*多个关键词匹配*/
func MultiMatch(w http.ResponseWriter, r *http.Request) {
	//过滤特殊字符
	vars := mux.Vars(r)

	q := elastic.NewMultiMatchQuery(vars["words"], "title", "text").Fuzziness("AUTO")
	searchResult, err := client.Search().
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
		Do(context.Background())

	if err != nil {
		panic(err)
	}

	util.JsonResponse(w, searchResult)

}

/*Add & update document*/
func NewDoc(doc *model.Doc) {
	d, err := client.Index().
		Index(conf.E.Index).
		Type(conf.E.Type).
		Id(doc.Id.String()).
		BodyJson(doc).
		Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
	}

	fmt.Printf("Indexed doc %s to index %s, type %s\n", d.Id, d.Index, d.Type)
}

/*Delete document*/
func DeleteDoc(id string) {

	res, err := client.Delete().
		Index(conf.E.Index).
		Type(conf.E.Type).
		Id(id).
		Do(context.TODO())

	if err != nil {
		panic(err)
	}

	_, err = client.Flush().Index(conf.E.Index).Do(context.TODO())
	if err != nil {
		panic(err)
	}

	fmt.Printf("\n Elastic %s  %s \n", id, res.Result)
}
