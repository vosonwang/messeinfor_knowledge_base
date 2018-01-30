package gjson

import (
	"testing"
	"io/ioutil"
	"fmt"
	"os/exec"
	"os"
	"strings"
	"reflect"
	"encoding/json"
	"github.com/tidwall/gjson"
)

type Base struct {
	Debug     bool
	Host      string `json:"Host,omitempty"`
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

type Conf struct {
	Base
	Pg
	Redis
}

var conf *Conf

func TestGjson(t *testing.T) {

	byteConf, err := ioutil.ReadFile("base.json")

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

	if base.Debug {
		if byteDev, err := ioutil.ReadFile("dev.json"); err != nil {
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
		if byteProd, err := ioutil.ReadFile("prod.json"); err != nil {
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

	conf=&Conf{base, pg, redis}

	fmt.Print(*conf)

	//fmt.Print(base)
	//fmt.Print(redis)
	//fmt.Print(pg)

	//byteBase, err := ioutil.ReadFile("base.conf.json")
	//
	//if err != nil {
	//	fmt.Print(err)
	//}
	//
	////fmt.Print(string(byteBase))
	//var base Base
	////
	//value := gjson.Get(string(byteBase), "base")
	//fmt.Print(value.Raw)
	//var str = `{"Protocol": "http://","Port": ":8300","ImagePath": "upload/images/","FilesPath": "upload/files/"}`
	//fmt.Print(str)
	//rs := gjson.GetBytes(byteBase, "base")
	//fmt.Print(rs.Raw)
	//fmt.Print(rs.Str)
	//fmt.Print([]byte(rs.Raw))
	//fmt.Print(byteBase)

	//json.Unmarshal([]byte(value.Raw), &base)
	////
	//fmt.Print(base)

	//byteJson, err := ioutil.ReadFile("conf.json")
	//
	//if err != nil {
	//	fmt.Print(err)
	//}
	//
	//base := gjson.GetBytes(byteJson, "base")
	////
	////
	////fmt.Print(typeof(base))
	//
	//m, ok := gjson.Parse(string(byteJson)).Value().(map[string]interface{})
	//if !ok {
	//	// not a map
	//
	//}
	//fmt.Print(m["base"])

	//if true {
	//	dev := gjson.GetBytes(byteJson, "dev")
	//	fmt.Print(dev)
	//} else {
	//	prod := gjson.GetBytes(byteJson, "prod")
	//	fmt.Print(prod)
	//}

	//os.Mkdir("test_go", 0777)

	//path := getCurrentPath()
	//fmt.Println(path)
	//fmt.Println(1)
	//json:=Load("conf.json")
	//value:=gjson.Get(string(*json), "base.store")
	//println(value.String())
	//fmt.Print(Load("../../dev.conf.json"))
}

func typeof(v interface{}) string {
	return reflect.TypeOf(v).String()
}

func Load(filename string) *[]byte {
	//dir, err := ioutil.ReadDir("/")
	//if err != nil {
	//	panic(err)
	//}
	//for _, fi := range dir {
	//	fmt.Print(fi)
	//}
	//
	//rd, err := ioutil.ReadDir("")
	//for _, fi := range rd {
	//	fmt.Println("")
	//	fmt.Println(fi.Name())
	//	fmt.Println(fi.IsDir())
	//	fmt.Println(fi.Size())
	//	fmt.Println(fi.ModTime())
	//	fmt.Println(fi.Mode())
	//}
	//fmt.Println("")
	//fmt.Println(err)

	data, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Print(err)
		return nil

	}
	p := []byte(data)
	return &p

}

func getCurrentPath() string {
	s, err := exec.LookPath(os.Args[0])
	checkErr(err)
	i := strings.LastIndex(s, "\\")
	path := string(s[0: i+1])
	return path
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
