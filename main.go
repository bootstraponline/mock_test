package main

import "fmt"

type HelloStruct struct {
}

func NewHelloStruct() *HelloStruct {
	return &HelloStruct{}
}

func (h *HelloStruct) Greeting() string {
	return "hello"
}

type GreetInterface interface {
	Greeting() string
}

func SayGreeting(g GreetInterface) {
	fmt.Println(g.Greeting())
}

func main() {
	hello := NewHelloStruct()

	SayGreeting(hello)
}