package main

import (
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"net/http"
	"time"
	"log"
	"messeinfor.com/messeinfor_knowledge_base/src/handler"
	"messeinfor.com/messeinfor_knowledge_base/src/conf"
	"messeinfor.com/messeinfor_knowledge_base/src/middleware"
	"messeinfor.com/messeinfor_knowledge_base/src/search"
)

func main() {

	/*Router*/
	r := mux.NewRouter()
	adminRouter := mux.NewRouter().PathPrefix("/admin").Subrouter().StrictSlash(true)

	/*-----普通权限路由：------*/
	/*搜索*/
	r.HandleFunc("/mkb/search/{words}", search.MultiMatch).Methods("GET")

	/*登录,请求token*/
	r.HandleFunc("/tokens", middleware.NewToken).Methods("POST")

	/*根据ID获取文档*/
	r.HandleFunc("/docs/{id}", handler.FindDoc).Methods("GET")

	/*根据ID获取别名*/
	r.HandleFunc("/alias/{id}", handler.FindAlias).Methods("GET")

	/*通过Alias和lang获取文档*/
	r.HandleFunc("/mkb/docAlias/{name}/{lang:[0-1]}", handler.FindDocByAlias).Methods("GET")

	/*获取图片*/
	r.HandleFunc("/upload/images/{name}", handler.GetImg).Methods("GET")
	//获取指定宽高的图片
	r.HandleFunc("/upload/images/{name}/{w:[0-9]+}/{h:[0-9]+}", handler.GetSizedImg).Methods("GET")
	//获取指定压缩比的图片
	r.HandleFunc("/upload/images/{name}/{percent:0.[0-9]+}", handler.GetPerceptualImg).Methods("GET")

	r.HandleFunc("/upload/files/{name}", handler.GetFile).Methods("GET")

	/*-----管理员权限路由：------*/

	/*别名*/
	//查询所有和描述接近的别名
	adminRouter.HandleFunc("/alias", handler.FindAliasByDesc).Methods("POST")

	/*别名与文章标题*/
	/*新增别名*/
	adminRouter.HandleFunc("/alias_titles", handler.NewAliasTitle).Methods("POST")
	/*获取所有别名（带中英文档标题）*/
	adminRouter.HandleFunc("/alias_titles", handler.FindAllAliasTitle).Methods("GET")
	/*删除别名和别名-文档关系*/
	adminRouter.HandleFunc("/alias_titles/{id}", handler.DeleteAliasTitle).Methods("DELETE")
	/*查找单条别名（带中英文档标题）*/
	adminRouter.HandleFunc("/alias_titles/{id}", handler.FindAliasTitle).Methods("GET")

	//根据语言查询和标题接近，并且未被占用的别名
	adminRouter.HandleFunc("/titles/{value}/{lang:[0-1]}", handler.FindTitle).Methods("GET")

	/*根据语言获取所有文档*/
	adminRouter.HandleFunc("/nodes/{lang:[0-1]}", handler.GetAllNodes).Methods("GET")
	/*添加文档和更新文档共用*/
	adminRouter.HandleFunc("/docs", handler.NewDoc).Methods("POST")
	//删除文档
	adminRouter.HandleFunc("/docs/{id}", handler.DeleteDoc).Methods("DELETE")

	//交换节点
	adminRouter.HandleFunc("/nodes/{down}/{up}", handler.SwapNode).Methods("PATCH")

	/*上传文件*/
	//上传图片
	adminRouter.HandleFunc("/images", handler.SaveImg).Methods("POST")
	//上传文件
	adminRouter.HandleFunc("/files/{name}", handler.SaveFile).Methods("POST")

	r.PathPrefix("/admin").Handler(negroni.New(
		negroni.HandlerFunc(middleware.ValidateToken),
		negroni.Wrap(adminRouter),
	))

	srv := &http.Server{
		Addr:    conf.B.Addr,
		Handler: r,
		// Good practice: enforce timeouts for servers you create!
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(srv.ListenAndServe())
}
