package main

import "fmt"

type animal interface {
	speak()
}

type animalSpeak struct {
	a animal
}

type cat struct{}

func (as *animalSpeak) animalSpeakFunc() {
	as.a.speak()
}

func (c *cat) speak() {
	fmt.Println("meow")
}

func main() {
	cat := &cat{}
	as := &animalSpeak{a: cat}
	as.animalSpeakFunc()
}
