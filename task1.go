package main

import (
	"fmt"
)

func is_vowel(char rune) bool {

	if (char == 'a') || (char == 'e') || (char == 'i') ||
		(char == 'o') || (char == 'u') {

		return true

	} else {

		return false

	}

}

func count_vowels(str string) string {
	var str1 string
	// fmt.Println("The Value of str", str)
	for _, char := range str {

		if is_vowel(char) {

			str1 += string(char)
			// fmt.Println("The value of str1", str1)
		}
	}

	return str1

}

func check(s string) string {
	m := make(map[rune]int)
	for _, r := range s {
		m[r]++
		// fmt.Println("The Value of m[r]", m[r])
	}
	for _, r := range s {
		//fmt.Println("rrer", r)
		if m[r] == 1 {
			//fmt.Println("The value of r", r)
			return string(r)
		}
	}
	return ""
}

func main() {

	x := count_vowels("india nation one")

	y := check(x)

	fmt.Println("The First Non Repeating Vowel is ", y)

}
