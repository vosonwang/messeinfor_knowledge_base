package handler

import (
	"net/http"
	"io/ioutil"
	"io"
	"time"
	"os"
	"messeinfor.com/messeinfor_knowledge_base/src/conf"
	"strings"
	"github.com/gorilla/mux"
	"fmt"
	"log"
	"github.com/disintegration/imaging"
	"strconv"
)

func SaveImg(w http.ResponseWriter, r *http.Request) {

	a := conf.B.ImagePath + time.Now().Format("2006-01-02_15-04-05")

	if f, err := os.Create(a); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "新建图片失败！")
	} else {
		defer f.Close()

		if _, err := io.Copy(f, r.Body); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "拷贝图像失败！")
		}

		JsonResponse(w, "/"+a)

	}

}

func GetImg(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	if imgByte, err := ioutil.ReadFile(conf.B.ImagePath + name); err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "读取图像失败！")
	} else {
		w.WriteHeader(http.StatusCreated)
		w.Write(imgByte)
	}

}

func GetSizedImg(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	width, _ := strconv.Atoi(vars["w"])
	height, _ := strconv.Atoi(vars["h"])

	if src, err := imaging.Open(conf.B.ImagePath + name); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "读取图像失败！")
	} else {
		dstImage128 := imaging.Resize(src, width, height, imaging.Lanczos)

		w.WriteHeader(http.StatusOK)
		imaging.Encode(w, dstImage128, imaging.PNG)
	}

}

func GetPerceptualImg(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	percent, _ := strconv.ParseFloat(vars["percent"], 32)

	if src, err := imaging.Open(conf.B.ImagePath + name); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "读取图像失败！")
	} else {

		dstImage128 := imaging.Resize(src, int(float64(src.Bounds().Dx())*percent), int(float64(src.Bounds().Dy())*percent), imaging.Lanczos)

		w.WriteHeader(http.StatusOK)
		imaging.Encode(w, dstImage128, imaging.PNG)
	}

}

func SaveFile(w http.ResponseWriter, r *http.Request) {
	//去除文件名中的空格，以便配合前端mavon能够正常显示链接
	fileName := strings.Replace(r.Header.Get("id"), " ", "", -1)
	a := conf.B.FilesPath + fileName

	var f *os.File

	defer f.Close()

	/*判断文件是否存在，如果存在则覆盖，没有则创建（没有考虑文件夹是否存在）*/
	if Exists(a) {
		if t, err := os.OpenFile(a, os.O_RDWR|os.O_CREATE, 0); err != nil {
			log.Print(err)
			fmt.Fprint(w, "打开文件失败！")
		} else {
			f = t
		}

	} else {
		if t, err := os.Create(a); err != nil {
			log.Print(err)
			fmt.Fprint(w, "创建文件失败！")
		} else {
			f = t
		}
	}

	if _, err := io.Copy(f, r.Body); err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "拷贝图像失败！")
	} else {
		JsonResponse(w, "/"+a)
	}

}

func GetFile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if file, err := ioutil.ReadFile(conf.B.FilesPath + id); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "获取文件失败！")
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write(file)
	}

}

/*
判断文件是否存在
 */
func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
