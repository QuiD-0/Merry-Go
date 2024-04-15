package main

import "fmt"

type EngineType string

const (
	Petrol   EngineType = "Petrol"
	Diesel   EngineType = "Diesel"
	Electric EngineType = "Electric"
)

type Car interface {
	Drive()
	Type()
}

type Tesla struct {
	Engine EngineType
}

func (t Tesla) Drive() {
	fmt.Println("drive")
}

func (t Tesla) Type() {
	fmt.Println("this is " + t.Engine)
}

func main() {

	var car Car = Tesla{Engine: Electric}
	car.Drive()
	car.Type()

}
