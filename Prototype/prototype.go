package main

import "fmt"

type shapeType int

const (
	circletype = 1
	squareType = 2
)

type shape interface {
	GetTypeID() shapeType
	printTypeProperty()
	clone() shape
}

type circle struct {
	ID            shapeType
	radius        int
	circumference float32
}

func NewCircleObject(radius int, circumference float32) *circle {
	return &circle{circletype, radius, circumference}
}

func (c *circle) clone() shape {
	return NewCircleObject(c.radius, c.circumference)
}

func (c *circle) GetTypeID() shapeType {
	return c.ID
}

func (c *circle) printTypeProperty() {
	fmt.Println("Circle", c.ID, c.radius, c.circumference)
}

func main() {
	circle := NewCircleObject(2, 88/7)

	ci := circle.(circle)
}
