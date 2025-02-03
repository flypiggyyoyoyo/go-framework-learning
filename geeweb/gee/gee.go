package gee

import (
	"fmt"
	"net/http"
)

//1+5  New+(GET+POST+Run+addRouter+ServeHttp)

type HanderFunc func(http.ResponseWriter, *http.Request)

type Engine struct {
	Router map[string]HanderFunc
}

func New() *Engine {
	return &Engine{make(map[string]HanderFunc)}
}

func (e *Engine) addRounter(method string, pattern string, h HanderFunc) {
	key := method + "-" + pattern
	e.Router[key] = h
}

func (e *Engine) GET(pattern string, h HanderFunc) {
	e.addRounter("GET", pattern, h)
}

func (e *Engine) POST(pattern string, h HanderFunc) {
	e.addRounter("POST", pattern, h)
}

func (e *Engine) Run(port string) (err error) {
	return http.ListenAndServe(port, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.Method + "-" + r.URL.Path
	if f, ok := e.Router[key]; ok {
		f(w, r)
	} else {
		fmt.Fprint(w, "404 not found")
	}
}
