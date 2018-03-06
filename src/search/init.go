package search

import (
	"context"

	"github.com/olivere/elastic"
	"messeinfor.com/messeinfor_knowledge_base/src/model"
	"log"
	"messeinfor.com/messeinfor_knowledge_base/src/conf"
	"fmt"
)

var client *elastic.Client


func init() {
	var err error

	client, err = elastic.NewClient(
		elastic.SetURL(conf.E.Url),
		elastic.SetSniff(false)) //停止嗅探，否则会报elastic node not found
	if err != nil {
		log.Print(conf.E.Url)
		panic(err)
	}

	// Use the IndexExists service to check if a specified index exists.
	// Starting with elastic.v5, you must pass a context to execute each service
	exists, err := client.IndexExists(conf.E.Index).Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
	}
	if !exists {
		mapping := `{
	"settings":{
		"number_of_shards": 1,
		"number_of_replicas": 0
	},
	"mappings":{
		"doc":{
			"properties":{
				"id":{
					"type":"text"
				},
				"number":{
					"type":"integer"
				},
				"lang":{
					"type":"keyword"
				},
				"title":{
					"type":"text",
					"boost":2,
            	    "analyzer":"ik_smart",
                	"search_analyzer":"ik_smart"
				},
				"text":{
					"type":"text",
            	    "analyzer":"ik_smart",
                	"search_analyzer":"ik_smart"
				},
				"alias_id":{
					"type":"text",
          			"doc_values": false
				},
				"parent_id":{
					"type":"text",
          			"doc_values": false
				},
				"creator":{
					"type":"text",
          			"doc_values": false
				},
				"updater":{
					"type":"text",
          			"doc_values": false
				},
				"created":{
					"type":"date"
				},
				"updated":{
					"type":"date"
				},
				"deleted":{
					"type":"date"
				}
			}
		}
	}
}`

		// Create a new index.
		createIndex, err := client.CreateIndex(conf.E.Index).BodyString(mapping).Do(context.Background())
		if err != nil {
			// Handle error
			panic(err)
		}
		if !createIndex.Acknowledged {
			// Not acknowledged
			panic(err)
		}

		importDoc()
	}

}

/*向ES中导入所有文档*/
func importDoc() {
	docs := model.FindAllDoc()
	if docs == nil {
		panic("获取不到文档")
	}

	for _, value := range *docs {
		NewDoc(&value)
	}

	var err error

	// Flush to make sure the documents got written.
	_, err = client.Flush().Index(conf.E.Index).Do(context.Background())
	if err != nil {
		panic(err)
	}

	count, err := client.Count(conf.E.Index).Do(context.TODO())
	if err != nil {
		panic(err)
	}

	fmt.Printf("\n Elastic:共索引 %d 个文档 \n",count)

}
