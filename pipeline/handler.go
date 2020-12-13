package pipeline

import (
	"github.com/Kutkh/constants"
)

type Handler struct {
	method         func(c KutkhData) Result
	parallelMethod []func(c KutkhData) Result
	context        KutkhData
	nextPipeline   *Handler
	err            error
}

type Result struct {
	result constants.Values
	err    error
}

type Param struct {
	param constants.Values
}

func (p *Handler) Next(f func(param KutkhData) Result) (res *Handler) {
	next := NewHandler()
	next.method = f
	p.nextPipeline = next
	return next
}

func (p *Handler) ParallelNext(functions ...func(c KutkhData) Result) (res *Handler) {
	next := NewHandler()
	next.parallelMethod = functions
	p.nextPipeline = next
	return next
}

func NewHandler() *Handler {
	p := &Handler{}
	p.context = KutkhData{}
	p.context.param = make(map[string]interface{})
	p.context.result = make(map[string]interface{})
	return p
}

func NewResult() Result {
	p := Result{}
	p.result = make(map[string]interface{})
	return p
}

func (self *Handler) SetValue(key string, value interface{}) {
	self.context.Param.param[key] = value
}

func (self *Result) SetResult(key string, value interface{}) {
	self.result[key] = value
}
