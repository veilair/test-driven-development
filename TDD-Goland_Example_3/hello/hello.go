package main

import "fmt"

func main() {
	fmt.Println(Hello("Rohan", "English"))
	fmt.Println(Hello("Raul", "Spanish"))
}

const helloPrefix = "Hello, "
const spanishPrefix = "Hola, "
const frenchPrefix = "Bonjour, "
const germanPrefix = "Hallo, "
const spanish = "Spanish"
const french = "French"
const german = "German"

//Hello function
func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := helloPrefix

	switch language {
	case spanish:
		return spanishPrefix + name

	case french:
		return frenchPrefix + name

	case german:
		return germanPrefix + name

	}

	return prefix + name

}
