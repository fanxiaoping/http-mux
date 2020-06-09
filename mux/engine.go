package mux

import (
	"net/http"
	"sync"
)

type Engine struct {
	tree map[string]HandlersChain
	Router
	pool sync.Pool
}

func NewEngine() *Engine{
	engine := &Engine{
		tree:        make(map[string]HandlersChain),
		Router: Router{
			Handlers:nil,
		},
	}
	engine.Router.engine = engine
	engine.pool.New = func() interface{} {
		return engine.allocateContext()
	}
	return engine
}

func (_self *Engine) allocateContext() *Context{
	return &Context{}
}

func (_self *Engine) ServeHTTP(w http.ResponseWriter,req *http.Request){
	c := _self.pool.Get().(*Context)
	c.Writer = w
	c.Request = req
	c.reset()

	_self.handleHTTPRequest(c)

	_self.pool.Put(c)
}

func (_self *Engine) handleHTTPRequest(c *Context){
	rPath := c.Request.URL.Path

	handlers := _self.getValue(rPath)
	if handlers != nil{
		c.handlers = handlers
		c.Next()
		return
	}
}

//获取路由下的相关HandlersChain
func (_self *Engine) getValue(path string)(handlers HandlersChain){
	handlers,ok := _self.tree[path]
	if !ok {
		return nil
	}
	return
}

func (_self *Engine) addRoute(path string,handlers HandlersChain){
	_self.tree[path] = handlers
}

func (_self *Engine) Use(middleware ...HandlerFunc){
	_self.Router.Use(middleware...)
}

