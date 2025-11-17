// crooruhe
package main

import (
	"fmt"
	"math/big"
	"strings"
	"unicode"
)

func Hexconvert(args string) interface{} {
	// this func is same as -> .SetString(hexNumber, 16)
	sum := new(big.Int)

	if strings.HasPrefix(args, "#") {
		args = args[1:]

		return Hashhexconvert(args)
	} else if strings.HasPrefix(args, "0x") {
		args = args[2:]
		length := len(args)
		counter := length - 1
		reversecounter := 0
		tempsum := new(big.Int)
		for counter > -1 {
			tempsum.SetInt64(0)
			if unicode.IsLetter(rune(args[counter])) {
				if args[counter]-55 > 15 {
					fmt.Println("Not a valid hexidecimal value")
					return nil
				}
				tempsum.Add(tempsum, new(big.Int).Mul(new(big.Int).SetInt64(int64(args[counter]-'0'-7)), new(big.Int).Exp(big.NewInt(16), big.NewInt(int64(reversecounter)), nil)))
			} else {
				tempsum.Add(tempsum, new(big.Int).Exp(big.NewInt(16), big.NewInt(int64(reversecounter)), nil))
				tempsum.Mul(tempsum, new(big.Int).SetInt64(int64(args[counter])-48))
			}
			sum.Add(sum, tempsum)
			counter--
			reversecounter++
		}
		return sum
	}
	fmt.Println("Invalid hexidecimal arguments")
	return nil
}
