package handler

import (
	"net/http"
	"messeinfor.com/messeinfor_knowledge_base/src/model"
	"github.com/gorilla/mux"
	"strconv"
	"fmt"
)

//func SwapNode(w http.ResponseWriter, r *http.Request) {
//	vars := mux.Vars(r)
//	id := vars["id"]
//	if down := model.FindAlias(id); down == nil {
//		w.WriteHeader(http.StatusInternalServerError)
//		fmt.Fprint(w, "数据库报错，找不到文档！")
//	} else {
//		var Map map[string]string
//		if err := json.NewDecoder(r.Body).Decode(&Map); err != nil {
//			w.WriteHeader(http.StatusInternalServerError)
//			fmt.Fprint(w, "无法解析节点")
//		} else {
//			if up := model.FindAlias(Map["id"]); up == nil {
//				w.WriteHeader(http.StatusInternalServerError)
//				fmt.Fprint(w, "数据库报错，找不到文档！")
//			} else {
//				if err := model.SwapNode(*down, *up); err {
//					JsonResponse(w, "交换成功！")
//				} else {
//					w.WriteHeader(http.StatusInternalServerError)
//					fmt.Fprint(w, "数据库报错，无法交换节点！")
//				}
//			}
//		}
//	}
//
//}

/*获取TOC所需要的所有节点*/
func GetAllNodes(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if lang, err := strconv.Atoi(vars["lang"]); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "语言参数不正确")
	} else {
		p := model.FindAllNodes(lang)
		JsonResponse(w, &p)
	}

}
