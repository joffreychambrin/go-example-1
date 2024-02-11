package main

import (
	"fmt"
)

func main() {
	constructors := map[string]func() Processor{
		"":    NewDefaultProcessor,
		"eks": NewEksProcessor,
	}

	p := &Pipeline{}
	p.Launch(getProcessor("", constructors))
	p.Launch(getProcessor("eks", constructors))
}

func getProcessor(name string, constructors map[string]func() Processor) Processor {
	// Check if constructor function exists for the given struct name
	constructor, ok := constructors[name]
	if !ok {
		fmt.Println("Constructor not found for", name)
		return nil
	}

	// Instantiate the struct using the constructor function
	processor := constructor()
	return processor
}
