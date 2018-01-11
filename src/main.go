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
	r.HandleFunc("/docs/{id}", handler.GetDoc).Methods("GET")

	r.HandleFunc("/upload/images/{id}", handler.GetImg).Methods("GET")
	r.HandleFunc("/upload/files/{id}", handler.GetFile).Methods("GET")

	/*获取所有文档和删除文档*/
	adminRouter.HandleFunc("/docs/{id}", handler.GetNodes).Methods("GET")
	adminRouter.HandleFunc("/docs/{id}", handler.DeleteDoc).Methods("DELETE")
	adminRouter.HandleFunc("/docs/{id}", handler.UpdateDoc).Methods("PATCH")
	/*创建Node和Doc都走这个路由*/
	adminRouter.HandleFunc("/docs", handler.NewDoc).Methods("POST")

	adminRouter.HandleFunc("/nodekey/{id}", handler.SwapNode).Methods("PATCH")

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
