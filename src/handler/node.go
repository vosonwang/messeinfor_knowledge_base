package handler

import (
	"net/http"
	"fmt"
	"encoding/json"
	"strconv"
	"messeinfor.com/messeinfor_knowledge_base/src/models"
	"github.com/gorilla/mux"
)

func SwapNode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	/*TODO 替换为FindNode*/
	if down, err := models.FindDoc(id); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "数据库报错，找不到文档！")
	} else {
		var Map map[string]string
		if err := json.NewDecoder(r.Body).Decode(&Map); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "无法解析节点")
		} else {
			if up, err := models.FindDoc(Map["id"]); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprint(w, "数据库报错，找不到文档！")
			} else {
				if err := models.Swap(down, up); err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					fmt.Fprint(w, "数据库报错，保存交换后的nodeKey失败！")
				} else {
					JsonResponse(w, "交换成功！")
				}
			}
		}
	}

}

func GetNodes(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if lang, err := strconv.Atoi(id); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "参数不正确")
	} else {
		if Nodes := models.FindNodes(lang); Nodes != nil {
			JsonResponse(w, Nodes)
		}
	}

}
