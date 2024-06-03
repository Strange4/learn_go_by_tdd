package main

import "fmt"

const (
	spanishLanguage    = "Spanish"
	frenchLanguage     = "French"
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Salut, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return fmt.Sprintf("%s%s", greetingPrefixSelector(language), name)
}

func greetingPrefixSelector(language string) (prefix string) {
	switch language {
	case spanishLanguage:
		prefix = spanishHelloPrefix
	case frenchLanguage:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
