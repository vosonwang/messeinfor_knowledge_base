package main

import (
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"net/http"
	"time"
	"log"
	"messeinfor.com/messeinfor_knowledge_base/src/conf"
	"messeinfor.com/messeinfor_knowledge_base/src/handler"
)

func main() {
	/*Router*/
	r := mux.NewRouter()
	adminRouter := mux.NewRouter().PathPrefix("/admin").Subrouter().StrictSlash(true)

	/*登录,请求token*/
	r.HandleFunc("/tokens", NewToken).Methods("POST")
	/*获取单个文档*/
	r.HandleFunc("/docs/{id}", handler.FindDoc).Methods("GET")

	r.HandleFunc("/upload/images/{id}", handler.GetImg).Methods("GET")
	r.HandleFunc("/upload/files/{id}", handler.GetFile).Methods("GET")

	/*获取所有文档和*/
	adminRouter.HandleFunc("/docs/{id}", handler.GetAllNodes).Methods("GET")
	/*添加文档*/
	adminRouter.HandleFunc("/docs", handler.AddNode).Methods("POST")
	//删除文档
	adminRouter.HandleFunc("/docs/{id}", handler.DeleteDoc).Methods("DELETE")
	//更新文档
	adminRouter.HandleFunc("/docs/{id}", handler.UpdateDoc).Methods("PUT")
	//交换节点
	adminRouter.HandleFunc("/alias/{id}", handler.SwapNode).Methods("PATCH")

	adminRouter.HandleFunc("/images", handler.SaveImg).Methods("POST")
	adminRouter.HandleFunc("/files", handler.SaveFile).Methods("POST")

	r.PathPrefix("/admin").Handler(negroni.New(
		negroni.HandlerFunc(ValidateToken),
		negroni.Wrap(adminRouter),
	))

	srv := &http.Server{
		Addr:    conf.WebPort,
		Handler: r,
		// Good practice: enforce timeouts for servers you create!
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(srv.ListenAndServe())
}
