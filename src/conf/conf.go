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
}

type Postgres struct {
	Host     string
	Port     int `json:"Port,int"`
	User     string
	Password string
	Db       string
}

type Redis struct {
	Host      string
	Port      string
	SecretKey string
}

var (
	B     Base
	R     Redis
	P     Postgres
	Debug = flag.Bool("debug", false, "-debug")
)

func init() {
	byteConf, err := ioutil.ReadFile("conf/base.json")

	if err != nil {
		fmt.Print(err)
	}

	err = json.Unmarshal(byteConf, &B)

	if err != nil {
		fmt.Print(err)
	}

	flag.Parse()

	if *Debug {
		if byteDev, err := ioutil.ReadFile("conf/dev.json"); err != nil {
			fmt.Print(err)
		} else {

			r := gjson.GetBytes(byteDev, "Redis")
			json.Unmarshal([]byte(r.Raw), &R)

			p := gjson.GetBytes(byteDev, "Pg")
			json.Unmarshal([]byte(p.Raw), &P)

		}

	} else {
		if byteProd, err := ioutil.ReadFile("conf/prod.json"); err != nil {
			fmt.Print(err)
		} else {

			r := gjson.GetBytes(byteProd, "Redis")
			json.Unmarshal([]byte(r.Raw), &R)

			p := gjson.GetBytes(byteProd, "Pg")
			json.Unmarshal([]byte(p.Raw), &P)

		}
	}
}
