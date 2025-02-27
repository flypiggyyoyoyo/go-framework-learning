package gee

import (
	"net/http"
	"strings"
)

//1+5  New+(GET+POST+Run+addRouter+ServeHttp)

type HandlerFunc func(*Context)

type (
	RouterGroup struct {
		prefix      string
		middlewares []HandlerFunc
		parent      *RouterGroup
		engine      *Engine
	}
	Engine struct {
		*RouterGroup
		Router *router
		Groups []*RouterGroup
	}
)

func New() *Engine {
	engine := &Engine{Router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.Groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

func (group *RouterGroup) Use(middlewares ...HandlerFunc) {
	group.middlewares = append(group.middlewares, middlewares...)
}

func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		parent: group,
		engine: engine,
	}
	engine.Groups = append(engine.Groups, newGroup)
	return newGroup
}

func (group *RouterGroup) addRouter(method string, comp string, h HandlerFunc) {
	pattern := group.prefix + comp
	group.engine.Router.addRoute(method, pattern, h)
}

func (group *RouterGroup) GET(pattern string, h HandlerFunc) {
	group.addRouter("GET", pattern, h)
}

func (group *RouterGroup) POST(pattern string, h HandlerFunc) {
	group.addRouter("POST", pattern, h)
}

func (e *Engine) Run(port string) (err error) {
	return http.ListenAndServe(port, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	middlewares := []HandlerFunc{}
	for _, group := range e.Groups {
		if strings.HasPrefix(req.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}
	c := NewContext(w, req)
	c.handlers = middlewares
	e.Router.handle(c)
}
