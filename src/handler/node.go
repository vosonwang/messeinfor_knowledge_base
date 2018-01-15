package handler

import (
	"net/http"
	"fmt"
	"strconv"
	"messeinfor.com/messeinfor_knowledge_base/src/model"
	"github.com/gorilla/mux"
	"log"
	"encoding/json"
)

func SwapNode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if down, err := model.FindAlias(id); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "数据库报错，找不到文档！")
	} else {
		var Map map[string]string
		if err := json.NewDecoder(r.Body).Decode(&Map); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "无法解析节点")
		} else {
			if up, err := model.FindAlias(Map["id"]); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprint(w, "数据库报错，找不到文档！")
			} else {
				if err := model.SwapNode(down, up); err != false {
					w.WriteHeader(http.StatusInternalServerError)
					fmt.Fprint(w, "数据库报错，无法交换节点！")
				} else {
					JsonResponse(w, "交换成功！")
				}
			}
		}
	}

}

/*获取TOC所需要的所有节点*/
func FindNodes(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if lang, err := strconv.Atoi(id); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "参数不正确")
	} else {
		if nodes, b := model.FindNodes(lang); b {
			JsonResponse(w, nodes)
		} else {
			log.Print(nodes)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "获取节点失败")
		}
	}

}
