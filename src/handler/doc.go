package handler

import (
	"net/http"
	"github.com/gorilla/mux"
	"messeinfor.com/messeinfor_knowledge_base/src/model"
	"fmt"
	"encoding/json"
	"github.com/satori/go.uuid"
)

func AddDoc(w http.ResponseWriter, r *http.Request) {
	var doc model.Doc
	if err := json.NewDecoder(r.Body).Decode(&doc); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "无法解析节点")
	}

	if Point:= model.AddDoc(doc); Point == nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "数据库报错，无法添加文档")
	} else {
		JsonResponse(w, *Point)
	}
}

func FindDoc(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if point := model.FindDoc(id); point == nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "数据库报错，找不到文档！")
	} else {
		JsonResponse(w, *point)
	}
}

func UpdateDoc(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var (
		doc model.Doc
		err error
	)

	if doc.Id, err = uuid.FromString(vars["id"]); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Id格式错误")
	}

	if err := json.NewDecoder(r.Body).Decode(&doc); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "无法解析节点")
	} else {
		if point := model.UpdateDoc(doc); point != nil {
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
