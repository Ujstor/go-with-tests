package main

import (
	"fmt"
)

const spanish = "Spanish"
const french = "French"
const engHelloPrefix = "Hello, "
const spanHelloPrefix = "Hola, "
const frHelloPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	} 

	if language == spanish {
		return spanHelloPrefix + name
	}

	if language == french {
		return frHelloPrefix + name
	}

	return engHelloPrefix + name 
}

func main() {
	fmt.Println(Hello("Chris", "")) }