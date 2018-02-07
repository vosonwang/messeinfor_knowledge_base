package cache

import (
	"github.com/go-redis/redis"
	"fmt"
	"messeinfor.com/messeinfor_knowledge_base/src/conf"
	"github.com/RedisLabs/redisearch-go/redisearch"
	"log"
	"messeinfor.com/messeinfor_knowledge_base/src/model"
	"time"
)

var client *redis.Client
var rSearch *redisearch.Client

func init() {
	/*Redis*/
	client = redis.NewClient(&redis.Options{
		Addr:     conf.R.Host + conf.R.Port,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	if pong, err := client.Ping().Result(); err != nil {
		fmt.Print(pong, err)
	}

	/*RediSearch*/
	rSearch = redisearch.NewClient(conf.R.Host+conf.R.Port, "mkb")

	a, _ := rSearch.Info()

	//判断索引是否已创建，如果已创建，则不创建新的索引，不清空数据
	if a == nil {
		// Create a schema
		sc := redisearch.NewSchema(redisearch.DefaultOptions).
			AddField(redisearch.NewTextFieldOptions("body", redisearch.TextFieldOptions{Weight: 4.0})).
			AddField(redisearch.NewTextFieldOptions("title", redisearch.TextFieldOptions{Weight: 5.0})).
			AddField(redisearch.NewNumericField("date"))

		// Drop an existing index. If the index does not exist an error is returned
		rSearch.Drop()

		// Create the index with the given schema
		if err := rSearch.CreateIndex(sc); err != nil {
			log.Fatal(err)
		}
	}

	ch := make(chan bool)
	go task(ch)

	if !<-ch {
		log.Print("创建RS索引失败！")
	}

	ticker := time.NewTicker(time.Second * 5).C
	go func() {
		for {
			<-ticker
			go task(ch)

			if !<-ch {
				log.Print("更新RS索引失败！")
			}
		}
	}()

}

/*定时任务
保证rs即使重启也能和数据库中的文档保持同步
*/
func task(b chan bool) {
	//查找rs中已经存在的文档$1
	a, err := client.Keys("????????-????-????-????-????????????").Result()
	if err != nil {
		log.Print(err)
		b <- false
	}

	//查找pg中所有的文档$2
	Ids := model.GetAllDocId()
	if Ids == nil {
		log.Print("model: 获取不到文档ID")
		b <- false
	}

	//找出数据库pg和缓存rs之间的差集
	//注意：这里必须要Ids在前
	add, del := DiffSlice(*Ids, a)

	//添加rs中没有的文档
	for _, v := range add {
		doc := model.FindDoc(v)

		if doc == nil {
			log.Print("model: 获取不到文档")
			b <- false
		}

		//将文档存入rSearch
		if !AddDoc(*doc) {
			log.Print("rSearch: 无法为文档建立索引")
			b <- false
		}
	}

	//删除rs中多余的文档
	for _, v := range del {
		_, err := client.Del(v).Result()

		if err != nil {
			log.Print("rSearch: 删除文档失败 " + v)
			b <- false
		}
	}

	b <- true
}

/*求两个字符串切片之间的差集*/
func DiffSlice(slice1, slice2 []string) (diffSlice1, diffSlice2 []string) {

	for _, v := range slice1 {
		if !InSlice(v, slice2) {
			diffSlice1 = append(diffSlice1, v)
		}
	}

	for _, v := range slice2 {
		if !InSlice(v, slice1) {
			diffSlice2 = append(diffSlice2, v)
		}
	}

	return
}

// InSlice checks given string in string slice or not.
func InSlice(v string, sl []string) bool {
	for _, vv := range sl {
		if vv == v {
			return true
		}
	}
	return false
}
