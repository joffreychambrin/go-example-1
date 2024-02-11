package main

type defaultProcessor struct {
}

func (d *defaultProcessor) preProcess() {
}

func (d *defaultProcessor) postProcess() {
}

func NewDefaultProcessor() Processor {
	return &defaultProcessor{}
}
