package decorator

/**
  a simple decorator
 */
import (
	"reflect"
)

type Base struct {
}

type Decorator interface {
	before()
	after()
}

func (self *Base) Decorator(decoPtr, fn interface{}) (err error) {
	var decoratedFunc, targetFunc reflect.Value
	decoratedFunc = reflect.ValueOf(decoPtr).Elem()
	targetFunc = reflect.ValueOf(fn)
	v := reflect.MakeFunc(targetFunc.Type(),
		func(in []reflect.Value) (out []reflect.Value) {
			self.before()
			out = targetFunc.Call(in)
			self.after()
			return
		})
	decoratedFunc.Set(v)
	return
}
