package main

import (
	"fmt"
)

func main() {
	// inputString := "A"
	// inputString := "ก"
	// inputString := "あ"

	fmt.Print("Enter the string to encode: ")
	var inputString string
	fmt.Scanln(&inputString)
	
	encodedHex := encodeToHex(inputString)
	fmt.Println("Encoded:", encodedHex)
}

func encodeToHex(input string) string {
	var encodedHex string
	var putPercent string
	var sumEncode string
	var formatFound bool = false
	var char_2_count int = 0

	for _, char_1 := range input {
		encodedHex = fmt.Sprintf("%X", char_1)
		for _, char_2 := range encodedHex {
			char_2_count++
			if (len(encodedHex) > 2 && !(formatFound)) {
				if (len(encodedHex) == 3) {
					putPercent += "%0" + string(char_2) + "%"
					formatFound = true
				} else if (len(encodedHex) == 4 && (char_2_count<2)) {
					putPercent += "%" + string(char_2)
				} else if (len(encodedHex) == 4 && (char_2_count==2)) {
					putPercent += string(char_2) + "%"
					formatFound = true
				}
			} else if (formatFound) {
				putPercent += string(char_2)
			} else {
				putPercent += "%00%" + string(char_2)
				formatFound = true
			}
		}
		sumEncode += putPercent
		putPercent = ""
		formatFound = false
		char_2_count = 0
	}
	return sumEncode
}