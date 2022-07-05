package main

import (
	"fmt"
	"time"
)

func print(c chan string) {
	for i := 0; i < 10; i++ {
		time.Sleep(100 * 100)
		fmt.Println("print..")

		c <- "kapil"
	}
}
func printer(c chan string) {
	for i := 0; i < 10; i++ {
		time.Sleep(500 * 1000)
		fmt.Println("i am in printer..")
		fmt.Println(<-c)

	}
}

func main() {

	c := make(chan string)

	go print(c)
	go printer(c)

	fmt.Println("main state waiting", <-c)
	var input string
	fmt.Scanln(&input)
	fmt.Println(input)

	in := "kulbhushan JIndal"
	ptr := &in
	const val = *ptr
	fmt.Println(in, ptr, val)
	//fmt.Println(input);
}
