package goee

import (
	"log"
	"net/http"
)

type HandlerFunc func(ctx *Context)

type Engine struct {
	router *router
	*RouterGroup
	groups []*RouterGroup // store all groups
}
func New() *Engine{
	//return &Engine{router: NewRouter()}
	engine := &Engine{router: NewRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine;
}

func (engine *Engine)addRouter(method,pattern string,handler HandlerFunc){
	engine.router.addRouter(method,pattern,handler);
}

func (engine *Engine) GET(pattern string,handler HandlerFunc){
	engine.addRouter("GET",pattern,handler)
}

func (engine *Engine) POST (pattern string,handler HandlerFunc){
	engine.addRouter("POST",pattern,handler)
}

func (engine *Engine) Run(addr string) error  {
	return http.ListenAndServe(addr,engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter,req *http.Request){
	c:= NewContext(w,req);
	engine.router.handle(c);
}

type RouterGroup struct {
	prefix string
	middlewares []HandlerFunc
	parent *RouterGroup
	engine *Engine
}
func (group *RouterGroup)Group(prefix string) *RouterGroup{
	engine := group.engine;
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		parent: group,
		engine: engine,
	}
	engine.groups = append(engine.groups,newGroup)
	return newGroup;
}
func (group *RouterGroup) addRouter(method string,comp string,handler HandlerFunc)  {
	pattern := group.prefix + comp;
	log.Printf("router %4s - %s",method,pattern)
	group.engine.addRouter(method,pattern,handler);
}
func (group *RouterGroup) GET(pattern string,handler HandlerFunc){
	group.addRouter("GET",pattern,handler)
}
func (group *RouterGroup) POST(pattern string,handle HandlerFunc){
	group.addRouter("POST",pattern,handle)
}