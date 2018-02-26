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

type Es struct {
	Url   string
	Index string
	Type  string
}

var (
	B     Base
	R     Redis
	P     Postgres
	E     Es
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

	var bs []byte

	if *Debug {
		bs, err = ioutil.ReadFile("conf/dev.json")
	} else {
		bs, err = ioutil.ReadFile("conf/prod.json")
	}

	if err != nil {
		fmt.Print(err)
	}

	r := gjson.GetBytes(bs, "Redis")
	json.Unmarshal([]byte(r.Raw), &R)

	p := gjson.GetBytes(bs, "Pg")
	json.Unmarshal([]byte(p.Raw), &P)

	e := gjson.GetBytes(bs, "es")
	json.Unmarshal([]byte(e.Raw), &E)

}
