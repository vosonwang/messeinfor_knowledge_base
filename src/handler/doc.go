package handler

import (
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
	"messeinfor.com/messeinfor_knowledge_base/src/models"
	"fmt"
	"github.com/satori/go.uuid"
)

func NewDoc(w http.ResponseWriter, r *http.Request) {
	if node := models.ParseNode(r.Body); node == nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "无法解析节点")
	} else {
		if id := models.AddDoc(node); id == uuid.Nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "数据库报错，无法添加文档")
		} else {
			JsonResponse(w, id)
		}
	}
}

func AllDoc(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if lang, err := strconv.Atoi(id); err != nil {

	} else {
		if docs := models.GetDocs(lang); docs != nil {
			JsonResponse(w, docs)
		}
	}
}
