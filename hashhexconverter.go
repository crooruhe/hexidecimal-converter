// crooruhe
package main

import (
	"fmt"
	"unicode"
)

func Hashhexconvert(hexcode string) []int {
	var result []int
	length := len(hexcode)
	if length%2 != 0 {
		fmt.Println("Not a hex number")
		return nil
	}
	for i := 0; i < length/2; i++ {
		idx := i * 2
		val := 0

		if unicode.IsLetter(rune(hexcode[idx])) {
			if hexcode[idx]-55 > 15 {
				fmt.Println("Not a valid hexidecimal value")
				return nil
			}
			val += (int(byte(hexcode[idx])) - 55) * 16
		} else {
			val += (int(hexcode[idx]) - 48) * 16
		}
		if unicode.IsLetter(rune(hexcode[idx+1])) {
			val += int(byte(hexcode[idx+1])) - 55
		} else {
			val += int(hexcode[idx+1]) - 48
		}

		result = append(result, val)
	}
	return result
}
