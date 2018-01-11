package handler

import (
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
	"messeinfor.com/messeinfor_knowledge_base/src/models"
	"fmt"
	"github.com/satori/go.uuid"
	"encoding/json"
)

func NewDoc(w http.ResponseWriter, r *http.Request) {
	var doc models.Doc
	if err := json.NewDecoder(r.Body).Decode(&doc); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "无法解析节点")
	}

	if id := models.AddDoc(doc); id == uuid.Nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "数据库报错，无法添加文档")
	} else {
		JsonResponse(w, id)
	}
}

func GetDoc(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if doc, err := models.FindDoc(id); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "数据库报错，找不到数据！")
	} else {
		JsonResponse(w, doc)
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

func UpdateDoc(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if doc, err := models.FindDoc(id); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "数据库报错，找不到文档！")
	} else {
		var newDoc models.Doc
		if err := json.NewDecoder(r.Body).Decode(&newDoc); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "无法解析节点")
		} else {
			doc.Text = newDoc.Text
			if updated, err := models.UpdateDoc(doc); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprint(w, "数据库更新文档失败")
			} else {
				JsonResponse(w, updated)
			}
		}

	}
}

func DeleteDoc(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if doc, err := models.FindDoc(id); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "数据库报错，找不到文档！")
	} else {
		if models.DeleteDoc(doc) == false {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "数据库报错，无法删除！")
		} else {
			JsonResponse(w, "删除成功")
		}
	}

}

func SwapNode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
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
