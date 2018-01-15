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

	if doc, err := model.AddDoc(doc); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "数据库报错，无法添加文档")
	} else {
		JsonResponse(w, doc)
	}
}

func FindDoc(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if doc, err := model.FindDoc(id); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "数据库报错，找不到文档！")
	} else {
		JsonResponse(w, doc)
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
		if doc, err := model.UpdateDoc(doc); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "数据库更新文档失败")
		} else {
			JsonResponse(w, doc)
		}
	}
}

func DeleteDoc(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if doc, err := model.FindDoc(id); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "数据库报错，找不到文档！")
	} else {
		if model.DeleteDoc(doc) == false {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "数据库报错，无法删除！")
		} else {
			JsonResponse(w, "删除成功")
		}
	}
}
