package main

import (
	"fmt"

	"github.com/enakamura3/go-plugin/greetings"
)

var ok bool

func main() {
	fmt.Println("go plugin sample")
	greetings.LoadUs()
	greetings.LoadBr()
}
