package main

const englishPrefix = "Hello, "
const spanishHelloPrefix = "Spanish, "
const frenchHelloPrefix = "French, "

func Hello(language string, name string) string {
	if name == "" {
		name = "Kyle"
	}

	return getPrefix(language) + name
}

func getPrefix(language string) string {
	switch language {
	case "french":
		return frenchHelloPrefix
	case "spanish":
		return spanishHelloPrefix
	default:
		return englishPrefix
	}
}
