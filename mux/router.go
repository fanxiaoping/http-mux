package mux

type Router struct {
	Handlers HandlersChain
	engine *Engine
}

func (_self *Router) Use(middleware ...HandlerFunc){
	_self.Handlers = append(_self.Handlers,middleware...)
}

func (_self *Router) AddRoute(absolutePath string,handlers ...HandlerFunc){
	handlers = _self.combineHandlers(handlers)
	//建立路由和相关中间件组的绑定
	_self.engine.addRoute(absolutePath,handlers)
}

//将定义的公用中间件和路由相关的中间件合并
func (_self *Router) combineHandlers(handlers HandlersChain) HandlersChain{
	finalSize := len(_self.Handlers) + len(handlers)
	mergedHandlers := make(HandlersChain,finalSize)
	copy(mergedHandlers,_self.Handlers)
	copy(mergedHandlers[len(_self.Handlers):],handlers)
	return mergedHandlers
}