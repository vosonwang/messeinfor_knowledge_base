package handler

import (
	"net/http"
	"github.com/gorilla/mux"
	"messeinfor.com/messeinfor_knowledge_base/src/model"
	"fmt"
	"encoding/json"
	"github.com/satori/go.uuid"
)

func AddNode(w http.ResponseWriter, r *http.Request) {
	var node model.Node
	if err := json.NewDecoder(r.Body).Decode(&node); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "无法解析节点")
	}

	if Point := model.AddNode(node); Point == nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "数据库报错，无法添加文档")
	} else {
		JsonResponse(w, *Point)
	}
}

func FindDoc(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if point := model.FindDocAlias(id); point == nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "数据库报错，找不到文档！")
	} else {
		JsonResponse(w, *point)
	}
}

func UpdateDoc(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var (
		docAlias model.DocAlias
		err      error
	)

	if docAlias.Id, err = uuid.FromString(vars["id"]); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Id格式错误")
	}

	if err := json.NewDecoder(r.Body).Decode(&docAlias); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "无法解析节点")
	} else {
		if point := model.UpdateDocAlias(docAlias); point == nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "数据库更新文档失败")
		} else {
			JsonResponse(w, *point)
		}
	}
}

func DeleteDoc(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if point := model.FindDoc(id); point == nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "数据库报错，找不到文档！")
	} else {
		if model.DeleteDoc(*point) == false {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "数据库报错，无法删除！")
		} else {
			JsonResponse(w, "删除成功")
		}
	}
}
