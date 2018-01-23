package handler

import (
	"net/http"
	"encoding/json"
	"fmt"
	"messeinfor.com/messeinfor_knowledge_base/src/model"
)

func AddAlias(w http.ResponseWriter, r *http.Request) {
	var alias model.Alias
	if err := json.NewDecoder(r.Body).Decode(&alias); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "无法解析节点")
	} else {
		if Point := model.NewAlias(alias); Point == nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "数据库:	无法添加文档")
		} else {
			JsonResponse(w, *Point)
		}
	}
}

func GetAllAlias(w http.ResponseWriter, r *http.Request) {
	if point := model.FindAllAlias(); point == nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "获取所有别名失败")
	} else {
		JsonResponse(w, *point)
	}
}
