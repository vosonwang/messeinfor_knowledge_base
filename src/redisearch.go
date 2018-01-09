package main

import (
	"fmt"
	"log"
	"time"

	"github.com/RedisLabs/redisearch-go/redisearch"
	"messeinfor.com/messeinfor_knowledge_base/src/conf"
)

func ExampleClient2() {

	// Create a client. By default a client is schemaless
	// unless a schema is provided when creating the index
	c := redisearch.NewClient(conf.Host+conf.Redis, "myIndex")

	// Create a schema
	sc := redisearch.NewSchema(redisearch.DefaultOptions).
		AddField(redisearch.NewTextField("body")).
		AddField(redisearch.NewTextFieldOptions("title", redisearch.TextFieldOptions{Weight: 5.0, Sortable: true})).
		AddField(redisearch.NewNumericField("date"))

	// Drop an existing index. If the index does not exist an error is returned
	c.Drop()

	// Create the index with the given schema
	if err := c.CreateIndex(sc); err != nil {
		log.Fatal(err)
	}

	// Create a document with an id and given score
	doc := redisearch.NewDocument("doc1", 1.0)
	doc.Set("title", "Redis").
		Set("body", "keep alive").
		Set("date", time.Now().Unix())

	// Index the document. The API accepts multiple documents at a time
	if err := c.IndexOptions(redisearch.DefaultIndexingOptions, doc); err != nil {
		log.Fatal(err)
	}

	// Searching with limit and sorting
	docs, total, err := c.Search(redisearch.NewQuery("alive").
		Highlight([]string{"alive"}, "<b>", "</b>").
		Limit(0, 2).
		SetReturnFields("title", "body"))

	fmt.Println(docs[0].Id, docs[0].Properties["title"], docs[0].Properties["body"], total, err)
	// Output: doc1 Hello world 1 <nil>
}
