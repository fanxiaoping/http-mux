package mux

import "net/http"

type Context struct {
	Request *http.Request
	Writer http.ResponseWriter
	handlers HandlersChain
	index int8
}

type HandlerFunc func(*Context)

type HandlersChain [] HandlerFunc

//模拟的调用堆栈
func (_self *Context) Next(){
	_self.index++
	for _self.index < int8(len(_self.handlers)){
		//按顺序执行HandlersChain内的函数
		//如果函数内无c.Next()方法调用则函数顺序执行完
		//如果函数内有c.Next()方法调用则代码执行到c.Next()方法处压栈，等待后面的函数执行完在回来执行c.Next()后的命令
		_self.handlers[_self.index](_self)
		_self.index++
	}
}

func (_self *Context) reset(){
	_self.handlers = nil
	_self.index = -1
}