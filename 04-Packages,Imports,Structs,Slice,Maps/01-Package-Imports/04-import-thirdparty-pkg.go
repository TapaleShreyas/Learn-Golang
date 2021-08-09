package mapper

import (
	"fmt"

	"github.com/rylans/getlang"
)

// use go get github.com/rylans/getlang to import third party package.

func main() {
	fmt.Println(Greet("Howdy, what's new?"))
	fmt.Println(Greet("Comment allez vous?"))
	fmt.Println(Greet("Wie geht es Ihnen?"))
}

func Greet(s string) string {
	info := getlang.FromString(s)
	switch info.LanguageCode() {
	case "en":
		return "Hello!"
	case "de":
		return "Guten Tag!"
	case "fr":
		return "Bonjour!"
	default:
		return "I don't know your language yet!"
	}
}
