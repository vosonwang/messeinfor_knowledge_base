package middleware

import (
	"github.com/go-redis/redis"
	"fmt"
	"messeinfor.com/messeinfor_knowledge_base/src/conf"
)

var rs *redis.Client

func init() {
	/*Redis*/
	rs = redis.NewClient(&redis.Options{
		Addr:     conf.R.Host + conf.R.Port,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := rs.Ping().Result()
	if err != nil {
		fmt.Print(pong, err)
	}

}
