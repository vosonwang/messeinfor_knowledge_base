package handler

import (
	"messeinfor.com/messeinfor_knowledge_base/src/model"
	"encoding/json"
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
)

func NewAliasTitle(w http.ResponseWriter, r *http.Request) {
	var aliasTitle model.AliasTitle
	if err := json.NewDecoder(r.Body).Decode(&aliasTitle); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "无法解析节点")
	} else {
		if Point := model.NewAliasTitle(aliasTitle); Point == nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "数据库:	无法添加别名记录")
		} else {
			JsonResponse(w, *Point)
		}
	}
}

func FindAllAliasTitle(w http.ResponseWriter, r *http.Request) {
	if p := model.FindAllAliasTitle(); p == nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "获取别名列表失败")
	} else {
		JsonResponse(w, *p)
	}
}

func DeleteAliasTitle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if id, err := uuid.FromString(vars["id"]); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "解析参数失败")
	} else {
		if model.DeleteAliasTitle(id) {
			JsonResponse(w, true)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "删除别名失败")
		}
	}

}

func FindAliasTitle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if p := model.FindAliasTitle(vars["id"]); p == nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "获取别名失败")
	} else {
		JsonResponse(w, &p)

	}
}
