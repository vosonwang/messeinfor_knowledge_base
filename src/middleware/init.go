package middleware

import (
	"github.com/go-redis/redis"
	"fmt"
	"messeinfor.com/messeinfor_knowledge_base/src/conf"
)

var client *redis.Client

func init() {
	/*Redis*/
	client = redis.NewClient(&redis.Options{
		Addr:     conf.R.Host + conf.R.Port,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Print(pong, err)
	}

}
