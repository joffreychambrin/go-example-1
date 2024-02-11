package main

import "fmt"

type eksProcessor struct {
}

func (t *eksProcessor) preProcess() {
	fmt.Println("preProcess")
}

func (t *eksProcessor) postProcess() {
	fmt.Println("postProcess")
}

func NewEksProcessor() Processor {
	return &eksProcessor{}
}
