package main

import (
	"fmt"
	"strconv"

	"github.com/Kutkh/decorator"
	"github.com/Kutkh/pipeline"
)

func main() {
	//build pipeline
	p := pipeline.NewStarter()
	// test method
	p.ParallelNext(goa, gob).Next(addNum).Next(toString).Next(printfString)
	//excute
	p.Execute()

	mybar := bar
	logDecorator := decorator.Log{}
	logDecorator.Decorator(&mybar, bar)
	mybar("hello,", "world!")

	fmt.Println()
}
func goa(ctx pipeline.Context) pipeline.Result {
	result := pipeline.NewResult()
	result.SetResult("a", 1)
	return result
}

func gob(ctx pipeline.Context) pipeline.Result {
	result := pipeline.NewResult()
	result.SetResult("b", 9)
	return result
}

func addNum(ctx pipeline.Context) pipeline.Result {
	//process
	a, _ := ctx.Value("a")
	b, _ := ctx.Value("b")
	c := a.(int) + b.(int)
	//build result
	result := pipeline.NewResult()
	result.SetResult("c", c)
	//return
	return result
}

func toString(ctx pipeline.Context) pipeline.Result {
	//process
	c, _ := ctx.Value("c")
	s := strconv.Itoa(c.(int))
	//build result
	result := pipeline.NewResult()
	result.SetResult("s", s)
	//return
	return result
}

func printfString(ctx pipeline.Context) pipeline.Result {
	fmt.Println("print")
	s, _ := ctx.Value("s")
	fmt.Println("output = " + s.(string))
	return pipeline.Result{}
}

func foo(a, b, c int) int {
	fmt.Printf("%d, %d, %d \n", a, b, c)
	return a + b + c
}
func bar(a, b string) string {
	fmt.Printf("%s, %s \n", a, b)
	return a + b
}
