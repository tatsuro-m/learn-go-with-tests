package main

import "fmt"

const (
	englishHelloPrefix = "Hello, "
	spanish            = "Spanish"
	spanishHelloPrefix = "Hola, "
	french             = "French"
	frenchHelloPrefix  = "Bonjour, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) string {
	var prefix string
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return prefix
}

func main() {
	fmt.Println(Hello("world", ""))
}
