package main

import (
	"fmt"
	"goeefframework/goee"
	"net/http"
)

func main(){
	//http.HandleFunc("/",indexHandler);
	//http.HandleFunc("/hello",helloHandler)
	//
	//err := http.ListenAndServe(":8080",nil)
	//if err != nil{
	//	fmt.Println(err)
	//	panic("err")
	//}
	//engine := new(goee.Engine)
	//http.ListenAndServe(":9999",engine)
	//字典书之前版本
	//engine := goee.New()
	//engine.Run(":9999")
	engine := goee.New()
	engine.GET("/", func(ctx *goee.Context) {
		ctx.HTMl(http.StatusOK,"<h1> Hello goee </h1>")
	})
}
func indexHandler(w http.ResponseWriter,req *http.Request){
	fmt.Fprintf(w,"url.path= %q\n",req.URL.Path)
}
func helloHandler(w http.ResponseWriter,req *http.Request){
	for k,v := range  req.Header{
		fmt.Fprintf(w,"header[%q]=%q \n",k,v)
	}
}