package main

import "fmt"

func main() {
	var hello = "Hello"
	var emojiString = "👋 🌍"
	concat := hello + " " + emojiString
	fmt.Println(concat)

	s2 := emojiString[:1]
	s3 := emojiString[2:]
	fmt.Println(emojiString, len(emojiString), s2, len(s2), s3, len(s3))

	// output
	// Hello 👋 🌍
	// 👋 🌍 9 � 1 �� 🌍 7

	// you will see special characters in the output
	// that because Go source code files are always written in utf-8(compact way to code unicode character).
	// when you use utf-8 it takes anyway from 1-4 bytes.
	// here each emoji takes 4 bytes.
	// as a result yo are getting 9 as length of emojiString
	// you can use rune (single literal) assign a literal and convert it to string

}
