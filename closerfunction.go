package main

import "fmt"

func adder() func(int)int{

	sum := 0

	return func(x int) int{

		return sum + x;
	}
}

func main(){

	closer1, closer2 := adder(), adder()

	fmt.Println("closer1 , closer2 " , closer1(4),closer2(4*6))
}