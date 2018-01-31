package conf

import (
	"io/ioutil"
	"fmt"
	"encoding/json"
	"github.com/tidwall/gjson"
	"flag"
)

type Base struct {
	Debug     bool
	Host      string `json:"Host,omitempty"`
	SecretKey string
	Protocol  string
	Port      string
	ImagePath string
	FilesPath string
}

type Pg struct {
	Host     string
	Port     int `json:"Port,int"`
	User     string
	Password string
	Db       string
}

type Redis struct {
	Host string
	Port string
}

type Config struct {
	Base
	Pg
	Redis
}

var (
	X     *Config
	Debug = flag.Bool("debug", false, "-debug")
)

func init() {
	byteConf, err := ioutil.ReadFile("src/conf/base.json")

	if err != nil {
		fmt.Print(err)
	}

	var (
		base  Base
		redis Redis
		pg    Pg
	)

	b := gjson.GetBytes(byteConf, "base")
	json.Unmarshal([]byte(b.Raw), &base)

	flag.Parse()


	if *Debug {
		if byteDev, err := ioutil.ReadFile("src/conf/dev.json"); err != nil {
			fmt.Print(err)
		} else {
			d := gjson.GetBytes(byteDev, "Dev.Host")
			json.Unmarshal([]byte(d.Raw), &base.Host)

			r := gjson.GetBytes(byteDev, "Dev.Redis")
			json.Unmarshal([]byte(r.Raw), &redis)

			p := gjson.GetBytes(byteDev, "Dev.Pg")
			json.Unmarshal([]byte(p.Raw), &pg)

		}

	} else {
		if byteProd, err := ioutil.ReadFile("src/conf/prod.json"); err != nil {
			fmt.Print(err)
		} else {
			d := gjson.GetBytes(byteProd, "Prod.Host")
			json.Unmarshal([]byte(d.Raw), &base.Host)

			r := gjson.GetBytes(byteProd, "Prod.Redis")
			json.Unmarshal([]byte(r.Raw), &redis)

			p := gjson.GetBytes(byteProd, "Prod.Pg")
			json.Unmarshal([]byte(p.Raw), &pg)

		}
	}

	X = &Config{base, pg, redis}
}
