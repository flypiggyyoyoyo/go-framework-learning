package gee

import (
	"net/http"
)

//1+5  New+(GET+POST+Run+addRouter+ServeHttp)

type HandlerFunc func(*Context)

type Engine struct {
	Router *Router
}

func New() *Engine {
	return &Engine{Router: NewRouter()}
}

func (e *Engine) addRouter(method string, pattern string, h HandlerFunc) {
	e.Router.addRoute(method, pattern, h)
}

func (e *Engine) GET(pattern string, h HandlerFunc) {
	e.addRouter("GET", pattern, h)
}

func (e *Engine) POST(pattern string, h HandlerFunc) {
	e.addRouter("POST", pattern, h)
}

func (e *Engine) Run(port string) (err error) {
	return http.ListenAndServe(port, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := NewContext(w, req)
	e.Router.handle(c)
}
