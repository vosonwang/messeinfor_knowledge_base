package handler

import (
	"net/http"
	"encoding/json"
	"fmt"
	"messeinfor.com/messeinfor_knowledge_base/src/model"
	"github.com/gorilla/mux"
	"messeinfor.com/messeinfor_knowledge_base/src/util"
)


func FindAlias(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	id := vars["id"]

	/*通过id查找*/
	if point := model.FindAlias(id); point == nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "找不到文档！")
	} else {
		util.JsonResponse(w, *point)
	}
}


//查询所有和描述接近的别名
func FindAliasByDesc(w http.ResponseWriter, r *http.Request) {
	a := make(map[string]string)
	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "无法解析节点")
	} else {
		if Point := model.FindAliasByDesc(a["description"]); Point == nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "数据库:	无法添加文档")
		} else {
			util.JsonResponse(w, *Point)
		}
	}
}
