package main

import "fmt"

// state
type binary interface {
	setTo()
}

// concrete
type one struct{}

func (o *one) setTo() {
	fmt.Println("Set to One!")
}

type zero struct{}

func (z *zero) setTo() {
	fmt.Println("Set to Zero!")
}

type context struct {
	b binary
}

func getBinaryContext() *context {
	return &context{
		b: &zero{},
	}
}

func (c *context) getContext() {
	c.b.setTo()
}

func (c *context) setContext(state binary) {
	c.b = state
}

func main() {
	binaryContext := getBinaryContext()
	binaryContext.getContext()
	binaryContext.setContext(&one{})
	binaryContext.getContext()
}
