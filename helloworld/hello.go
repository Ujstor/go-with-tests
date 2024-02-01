package main

import (
	"fmt"
)

const (
	spanish = "Spanish"
	french = "French"
	hrvatski = "Cro"

	engHelloPrefix = "Hello, "
	spanHelloPrefix = "Hola, "
	frHelloPrefix = "Bonjour, "
	croHelloPrefix = "Bok, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	} 
	
	return grettingPrefix(language) + name
}

func grettingPrefix(language string) (prefix string) {

	switch language {
	case french: 
		prefix = frHelloPrefix
	case spanish:
		prefix = spanHelloPrefix
	case hrvatski:
		prefix = croHelloPrefix
	default:
		prefix = engHelloPrefix
	}

	return prefix
}

func main() {
	fmt.Println(Hello("Ivan", "")) }