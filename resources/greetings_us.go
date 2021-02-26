package main

import "fmt"

type US struct{}

var Us US

func (us US) Hello() {
	fmt.Println("Hello!")
}

func HelloUS() {
	fmt.Println("Hello!")
}
