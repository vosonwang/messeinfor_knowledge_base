package handler

import (
	"net/http"
	"messeinfor.com/messeinfor_knowledge_base/src/model"
	"github.com/gorilla/mux"
	"strconv"
	"fmt"
)

/*获取TOC所需要的所有节点*/
func GetAllNodes(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	lang, _ := strconv.Atoi(vars["lang"])

	p := model.FindAllNodes(lang)
	JsonResponse(w, &p)

}

func SwapNode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if err := model.SwapNode(vars["down"], vars["up"]); err {
		JsonResponse(w, "交换成功！")
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "数据库报错，无法交换节点！")
	}

}
