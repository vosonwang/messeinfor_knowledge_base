package handler

import (
	"net/http"
	"messeinfor.com/messeinfor_knowledge_base/src/model"
	"encoding/json"
	"fmt"
)

func FindTitle(w http.ResponseWriter, r *http.Request) {
	var title model.Title
	if err := json.NewDecoder(r.Body).Decode(&title); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "无法解析节点")
	} else {
		if Point := model.FindTitles(title); Point == nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "数据库:	无法添加文档")
		} else {
			JsonResponse(w, *Point)
		}
	}
}
