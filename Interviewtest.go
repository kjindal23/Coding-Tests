package main

import ("fmt"
"time"
)

func pinger(c chan string){
	for i:=0; ; i++{
		c<- "ping number of times"
	}
}
func printer(c chan string){
	for {
		msg := <- c
		fmt.Println("print from printer", msg)
		time.Sleep(time.Second * 1)
	}
}

func main(){
	var c chan string = make(chan string)
	go pinger(c)
	go printer(c)

	fmt.Println(<-c)

	//var input string
	//fmt.Scanln(&input)
}