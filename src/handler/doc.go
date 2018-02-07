package handler

import (
	"net/http"
	"github.com/gorilla/mux"
	"messeinfor.com/messeinfor_knowledge_base/src/model"
	"fmt"
	"encoding/json"
	"strconv"
)

func NewDoc(w http.ResponseWriter, r *http.Request) {
	var doc model.Doc
	if err := json.NewDecoder(r.Body).Decode(&doc); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "无法解析节点")
	} else {
		if p := model.NewDoc(doc); p != nil {
			JsonResponse(w, *p)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "数据库:	无法添加文档")
		}
	}
}



func FindDoc(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	/*通过id查找*/
	if point := model.FindDoc(id); point == nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "找不到文档！")
	} else {
		JsonResponse(w, *point)
	}

}

func DeleteDoc(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if point := model.FindDoc(id); point == nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "数据库：找不到文档！")
	} else {
		if model.DeleteDoc(*point) == false {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "数据库：无法删除！")
		} else {
			JsonResponse(w, "删除成功")
		}
	}
}

func FindDocByAlias(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	lang, _ := strconv.Atoi(vars["lang"])

	if p := model.FindAliasByName(name); p == nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "找不到别名")
	} else {
		if point := model.FindDocByAlias(p.Id, lang); point == nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "找不到文档")
		} else {
			JsonResponse(w, *point)
		}
	}

}
