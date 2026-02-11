package main

import "fmt"

func main() {
	// The for-range value is a copy
	evenVals := []int{2, 4, 6, 8, 10}
	for _, v := range evenVals {
		v += 1
	}
	fmt.Println(evenVals) // Output: [2, 4, 6, 8, 10]

	// Switch statement
	words := []string{"a", "cow", "smile", "gopher"}
	for _, word := range words {
		switch size := len(word); size {
		case 0, 1, 2, 3, 4:
			fmt.Println(word, " is a small word")
		case 5:
			wordLen := len(word)
			fmt.Println(word, " is exactly the right size ", wordLen)
		case 6, 7, 8, 9, 10:
		default:
			fmt.Println(word, " is a giant word")
		}
	}

	// Blank switch
	for _, word := range words {
		switch size := len(word); {
		case size < 5:
			fmt.Println(word, " is a small word")
		case size == 5:
			wordLen := len(word)
			fmt.Println(word, " is exactly the right size ", wordLen)
		case size > 10:
			fmt.Println(word, " is a giant word")
		}
	}

	// Or even more concise
	for _, word := range words {
		switch {
		case len(word) < 5:
			fmt.Println(word, " is a small word")
		case len(word) == 5:
			fmt.Println(word, " is exactly the right size ", len(word))
		case len(word) > 10:
			fmt.Println(word, " is a giant word")
		}
	}
}
