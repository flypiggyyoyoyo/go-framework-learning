package gee

import "log"

type Router struct {
	handler map[string]HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		handler: make(map[string]HandlerFunc),
	}
}

func (r *Router) addRoute(Method string, Pattern string, Handler HandlerFunc) {
	log.Printf("Route %4s - %s", Method, Pattern)
	key := Method + "-" + Pattern
	r.handler[key] = Handler
}

func (r *Router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if h, ok := r.handler[key]; ok {
		h(c)
	} else {
		c.String(404, "404", c.Path)
	}
}
