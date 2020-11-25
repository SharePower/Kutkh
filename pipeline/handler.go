package pipeline

import (
	"awesomeProject/util"
)


type Handler struct {
	method         func(c Context) Result
	parallelMethod []func(c Context) Result
	context        Context
	nextPipeline   *Handler
	err error
}


type  Result struct {
	result util.Values
	err  error
}

type  Param struct {
	param util.Values
}

func (p *Handler) Next(f func(param Context) Result)(res *Handler){
	next := NewHandler()
	next.method = f
	p.nextPipeline = next
	return next
}

func  (p *Handler) ParallelNext(functions ... func(c Context) Result)(res *Handler) {
	next := NewHandler()
	next.parallelMethod = functions
	p.nextPipeline = next
	return next
}

func NewHandler() *Handler {
	p :=  &Handler{}
	p.context = Context{}
	p.context.param = make(map[string]interface{})
	p.context.result = make(map[string]interface{})
	return p
}

func NewResult() Result {
	p := Result{}
	p.result = make(map[string]interface{})
	return p
}



func (self *Handler) SetValue(key string, value interface{})  {
	self.context.Param.param[key] = value
}

func (self *Result) SetResult(key string, value interface{})  {
	self.result[key] = value
}