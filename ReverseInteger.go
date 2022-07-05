package main

import (
	"fmt"
	"strconv"
)

/*func reverse(value int) int {

	number := 0
	for value != 0 {
		reminder := value % 10
		number = number * 10
		number = number + reminder
		value = value / 10
	}
	return number
}*/

func reversestring(strvalue string) string {

	runes := []rune(strvalue)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)

}

/*
func reversestring1(strvalue string) string {

	ptr := &strvalue
	lenght := len(strvalue)

	fmt.Println("The value of ptr and len = ", (ptr),lenght);

	fmt.Println("The value of ptr and len = ", (ptr++),lenght);

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)

}*/

func main() {

	intValue := 1234
	fmt.Println("the value of integer == ", intValue)
	strvalue := strconv.Itoa(intValue)
	reversenumber := reversestring(strvalue)
	reverseintvalue, _ := strconv.Atoi(reversenumber)

	fmt.Println("reverse number = ", reverseintvalue)

}
