package main

import (
	"github.com/a3no/go-plugin-example/plugin"
)

func main() {
	plugin.Serve(new(GreeterHello))
}


// Here is a real implementation of Greeter
type GreeterHello struct{}

func (GreeterHello) Greet() string { return "Hello!Hello!!" }
