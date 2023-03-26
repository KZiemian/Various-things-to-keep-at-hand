package main

import "fmt"

func checkVowels(s string) {
	n := len(s)
	positionOfA := -1

	if s[0] == 'a' {
		positionOfA = 0
	} 

	for i := 1; i < n; {
		if s[i] == 'e' {
			i++
		} else if s[i] == 'a' {
			positionOfA = i
			i++

			continue
		} else {
			positionOfA = -1
			i++

			continue
		}

		if s[i] == 'i' {
			i++
		} else if s[i] == 'a' {
			positionOfA = i
			i++

			continue
		} else {
			positionOfA = -1
			i++

			continue
		}

		if s[i] == 'o' {
			i++
		} else if s[i] == 'a' {
			positionOfA = i

			continue
		} else {
			positionOfA = -1
			i++
 
			continue
		}

		if s[i] == 'u' {
			fmt.Printf("String \"aeiou\" znaleziony na pozycje %v w stringu: \"%v\".\n", positionOfA, s)

			return 
		} else if s[i] == 'a' {
			positionOfA = i
			i++

			continue
		} else {
			positionOfA = -1
			i++

			continue
		}
		 
	}

	fmt.Printf("String \"aeiou\" nie jest zawarty w stringu: %v.\n", s)
}

func main() {
	strVar := "aeiou"

	checkVowels(strVar)

	strVar = "aaeiou"

	checkVowels(strVar)

	strVar = "caeiou"

	checkVowels(strVar)

	strVar = "aaaeiou"

	checkVowels(strVar)

	strVar = "aeiouu"

	checkVowels(strVar)

	strVar = "abcdaeiouefg"

	checkVowels(strVar)

	strVar = "abcdefg"

	checkVowels(strVar)

	strVar = "aeiiou"

	checkVowels(strVar)
}
