package pipeline

type  Context struct {
	Param
	Result
	err  error
}


func (self *Context) DefaultValue(key string, defaultValue interface{}, valueType string) (interface{}, bool) {
	values,err := self.Param.param[key]
	return values,err
}

func (self *Context) Value(key string) (interface{},bool) {
	values,err := self.Param.param[key]
	return values,err
}

func (self *Context) SetValue(key string, value interface{})  {
	self.Param.param[key] = value
}