package main

import "fmt"

func main(){

	teststring := "kulbhushan Jindal";
	var reverse string;
	for _, value := range teststring {
		reverse = string(value) + reverse
	}

	fmt.Println("reverse string == ", reverse)
}