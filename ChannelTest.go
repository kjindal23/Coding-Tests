package main

import (
	"fmt"
)

func countNumber(c chan int, num int) {
	for i := 0; i < num; i++ {
		fmt.Println("value of i=", i)

	}
	fmt.Println("the value is = ", <-c)
	c <- 2
	//return i+i
}

func countNumber1(c chan string, num int) {
	for i := 0; i < num; i++ {
		fmt.Println("value of i=", i)

	}
	//fmt.Println("the value is = ", <-c)
	//c <- "done2"
	//return i+i
}

func main() {

	//c := make(chan string, 2)
	test := make(chan int)
	//test <- 3
	//fmt.Println(<-test)
	//test <- 5
	//fmt.Println(test)
	//count(5)
	go countNumber(test, 5)
	//fmt.Println("the value is = ", <-c)
	fmt.Println("the value is = ", <-test)
     
	//go countNumber1(c, 4)

	//fmt.Println("the value is = ", c)
	//fmt.Println("the value is = ", <-c)

}
