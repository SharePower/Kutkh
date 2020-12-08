package pipeline

type KutkhData struct {
	Param
	Result
	err error
}

func (self *KutkhData) DefaultValue(key string, defaultValue interface{}, valueType string) (interface{}, bool) {
	values, err := self.Param.param[key]
	return values, err
}

func (self *KutkhData) Value(key string) (interface{}, bool) {
	values, err := self.Param.param[key]
	return values, err
}

func (self *KutkhData) SetValue(key string, value interface{}) {
	self.Param.param[key] = value
}
