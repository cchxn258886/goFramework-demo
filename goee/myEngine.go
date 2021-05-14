package goee

import (
	"fmt"
	"net/http"
)
type HandlerFuncA func(w http.ResponseWriter,r *http.Request)
//type HandlerFuncA func(ctx *Context)
type EngineA struct {
	router map[string] HandlerFuncA;
};

func NewA() *EngineA{
	return &EngineA{router: make(map[string]HandlerFuncA)}
}
//func (engine *Engine) ServeHTTP(w http.ResponseWriter,req *http.Request){
//	switch req.URL.Path {
//	case "/":
//		fmt.Fprintf(w,"url.path = %q\n",req.URL.Path)
//	case "/hello":
//		for k,v := range req.Header{
//			fmt.Fprintf(w,"header[%q] = %q \n",k,v)
//		}
//	default:
//		fmt.Fprintf(w,"404 not found %s \n",req.URL)
//	}
//}
func (engine *EngineA) ServeHTTP(w http.ResponseWriter,req *http.Request){
	key := req.Method + "-" + req.URL.Path;
	if handle,ok := engine.router[key];ok{
		handle(w,req);
	}else {
		fmt.Fprintf(w,"404 not found :%s \n",req.URL)
	}
	//c := NewContext(w,req)
}

func (engine *EngineA) addRoute(method string,pattern string,handle HandlerFuncA){
	key := method+ "-"+pattern;
	engine.router[key] = handle;
}
func (engine *EngineA) Get(pattern string,handler HandlerFuncA){
	engine.addRoute("Get",pattern,handler);
}

func (enging *EngineA) POST(pattern string,handler HandlerFuncA){
	enging.addRoute("POST",pattern,handler)
}
func (engine *EngineA) Run(address string) (err error) {
	return http.ListenAndServe(address, engine)
}