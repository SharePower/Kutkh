package decorator

import (
	"fmt"
)

type Log struct {
	 Base
}

func (self *Log) before() {
	fmt.Println(1)
}

func (self *Log) after() {
	fmt.Println(2)
}

