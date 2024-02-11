package main

import "fmt"

type Processor interface {
	preProcess()
	postProcess()
}

type Pipeline struct {
}

func (p *Pipeline) Launch(processor Processor) {
	fmt.Println("in launch method of pipeline")

	fmt.Println("before preProcess")
	processor.preProcess()
	fmt.Println("after preProcess")

	fmt.Println("before postProcess")
	processor.postProcess()
	fmt.Println("after postProcess")
}
