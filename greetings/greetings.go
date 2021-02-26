package greetings

import (
	"os"
	"path/filepath"
	"plugin"
)

type Greetings interface {
	Hello()
}

var BASE_PATH string

func LoadUs() {
	loadBasepath()
	filePath := filepath.Join(BASE_PATH, "resources", "greetings_us.so")
	p, err := plugin.Open(filePath)
	if err != nil {
		panic(err)
	}

	// Look for a type named "Us"
	t, err := p.Lookup("Us")
	if err != nil {
		panic(err)
	}

	// Type assertion to Greeting interface, since implements Hello method
	g, ok := t.(Greetings)
	if !ok {
		panic("type assertion fail!")
	}

	// Execute the implemented method
	g.Hello()

	// Look for a function named "HelloUS"
	f, err := p.Lookup("HelloUS")
	if err != nil {
		panic(err)
	}

	// Execute the function
	f.(func())()
}

func LoadBr() {
	loadBasepath()
	filePath := filepath.Join(BASE_PATH, "resources", "greetings_br.so")
	p, err := plugin.Open(filePath)
	if err != nil {
		panic(err)
	}

	// Look for a type named "Bs"
	t, err := p.Lookup("Br")
	if err != nil {
		panic(err)
	}

	// Type assertion to Greeting interface, since implements Hello method
	g, ok := t.(Greetings)
	if !ok {
		panic("type assertion fail!")
	}

	// Execute the implemented method
	g.Hello()

	f, err := p.Lookup("HelloBR")
	if err != nil {
		panic(err)
	}

	// Execute the function
	f.(func())()
}

// Get working directory path
func loadBasepath() {
	if len(BASE_PATH) == 0 {
		BASE_PATH, _ = os.Getwd()
	}
}
