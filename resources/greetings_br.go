package main

import "fmt"

type BR struct{}

var Br BR

func (br BR) Hello() {
	fmt.Println("Olá!")
}

func HelloBR() {
	fmt.Println("Olá!")
}
