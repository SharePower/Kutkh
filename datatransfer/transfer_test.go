package datatransfer

import (
	"fmt"
	"strings"
	"testing"
)

type CatData struct {
	Model
	Name string
}

type DogData struct {
	Model
	Name string
}

type DogToCatTransferRule struct {
	TransferRule
}

func (self *DogToCatTransferRule) execute(source Transferable, target Transferable) {
	// this is a simple demo to describe this func means
	dog := source.(*DogData)
	cat := target.(*CatData)

	name := dog.Name
	prefix := strings.Split(name, ",")[0]
	cat.Name = prefix + ", cat"
	return
}

func TestTransfer(t *testing.T) {

	dog := &DogData{}
	cat := &CatData{}
	rule := &DogToCatTransferRule{}

	dog.Name = "hello, dog"

	Transfer(dog, cat, rule)

	exceptName := "hello, cat"
	factName := cat.Name
	fmt.Println(exceptName == factName)
}
