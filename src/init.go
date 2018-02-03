package main

import (
	"log"
	"github.com/go-redis/redis"
	"messeinfor.com/messeinfor_knowledge_base/src/conf"
	"fmt"
)

var client *redis.Client



func init() {
	/*Log*/
	log.SetPrefix("TRACE: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)


	/*Redis*/
	client = redis.NewClient(&redis.Options{
		Addr:      conf.R.Host+ conf.R.Port,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	if pong, err := client.Ping().Result(); err != nil {
		fmt.Print(pong, err)
	}

}
