package conf

import (
	"io/ioutil"
	"fmt"
	"encoding/json"
	"github.com/tidwall/gjson"
	"flag"
)

type Base struct {
	Addr      string
	ImagePath string
	FilesPath string
	SecretKey string `json:"SecretKey,omitempty"`
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
	byteConf, err := ioutil.ReadFile("conf/base.json")

	if err != nil {
		fmt.Print(err)
	}

	var (
		base  Base
		redis Redis
		pg    Pg
	)

	err = json.Unmarshal(byteConf, &base)

	if err != nil {
		fmt.Print(err)
	}

	flag.Parse()

	if *Debug {
		if byteDev, err := ioutil.ReadFile("conf/dev.json"); err != nil {
			fmt.Print(err)
		} else {

			s := gjson.GetBytes(byteDev, "SecretKey")
			json.Unmarshal([]byte(s.Raw), &base.SecretKey)

			r := gjson.GetBytes(byteDev, "Redis")
			json.Unmarshal([]byte(r.Raw), &redis)

			p := gjson.GetBytes(byteDev, "Pg")
			json.Unmarshal([]byte(p.Raw), &pg)

		}

	} else {
		if byteProd, err := ioutil.ReadFile("conf/prod.json"); err != nil {
			fmt.Print(err)
		} else {

			s := gjson.GetBytes(byteProd, "SecretKey")
			json.Unmarshal([]byte(s.Raw), &base.SecretKey)

			r := gjson.GetBytes(byteProd, "Redis")
			json.Unmarshal([]byte(r.Raw), &redis)

			p := gjson.GetBytes(byteProd, "Pg")
			json.Unmarshal([]byte(p.Raw), &pg)

		}
	}

	X = &Config{base, pg, redis}
}
