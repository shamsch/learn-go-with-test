package main

import "fmt"

const (
	spanish = "Spanish"
	french  = "French"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

// here we have named prefix which is NAMED RETURN which means a variable is by that name is accessible with default value to 0, "", based on type
// it doesn't need to be mentioned explicitly in return 
func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}

// public function start with uppercase i.e Hello() - it is accessible from all files in the modules
// private function start with lowercase i.e greetingPrefix()- not accessible from other files in the modules