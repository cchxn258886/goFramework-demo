package goee

import (
	"fmt"
	"net/http"
	"strings"
)

//用roots存储每种请求方式的字典树根节点 使用handlers存储HandleFunc
type router struct {
	//字典树
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

func NewRouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc)}
}

//字典树
func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")
	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}

func (r *router) addRouter(method string, pattern string, handler HandlerFunc) {
	//字典树
	parts := parsePattern(pattern)

	fmt.Printf("router %4s - %s \n", method, pattern)
	key := method + "-" + pattern
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}
	r.roots[method].insert(pattern, parts, 0)
	r.handlers[key] = handler
	//r.handlers[key] = handler
}
func (r *router) getRoute(method string, path string) (*node, map[string]string) {
	searchParts := parsePattern(path)
	params := make(map[string]string)
	root, ok := r.roots[method]
	n := root.search(searchParts, 0)
	if !ok {
		return nil, nil
	}
	if n != nil {
		parts := parsePattern(n.pattern)
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[index]
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		return n, params
	}
	return nil, nil
}

func (r *router) handle(c *Context) {
	//字典树
	n, params := r.getRoute(c.Method, c.Path)
	if n != nil {
		c.Params = params
		key := c.Method + "-" + n.pattern
		r.handlers[key](c)
	} else {
		c.String(http.StatusNotFound, "404 not found", c.Path)
	}
	//老版本代码
	//key := c.Method + "-" + c.Path
	//if handler,ok := r.handlers[key];ok{
	//	handler(c)
	//}else {
	//	c.String(http.StatusNotFound,"404 not found:%s \n",c.Path)
	//}
}
