package goee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

type Context struct {
	//必须存在的info
	req       *http.Request
	respWrite http.ResponseWriter
	//request info
	Path   string
	Method string
	//字典树参数
	Params map[string]string
	//resp info
	StatusCode int
	//middleware
	handlers []HandlerFunc
	index    int
}

//字典树
func (c *Context) Param(key string) string {
	value, _ := c.Params[key]
	return value
}

func NewContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		req: req, respWrite: w, Path: req.URL.Path, Method: req.Method, index: -1,
	}
}

func (c *Context) PostForm(key string) string {
	return c.req.FormValue(key)
}

func (c *Context) Query(key string) string {
	//println()
	return c.req.URL.Query().Get(key)
}

func (c *Context) Status(code int) {
	c.StatusCode = code
	c.respWrite.WriteHeader(code)
}

func (c *Context) SetHeader(key string, value string) {
	c.respWrite.Header().Set(key, value)
}

func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.respWrite.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) Json(code int, obj interface{}) {
	c.SetHeader("Context-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.respWrite)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.respWrite, err.Error(), 500)
	}
}

func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.respWrite.Write(data)
}
func (c *Context) HTMl(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.respWrite.Write([]byte(html))
}
func (c *Context) NEXT() {
	c.index++
	s := len(c.handlers)
	for ; c.index <s;c.index++{
		c.handlers[c.index](c)
	}
}
