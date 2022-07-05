package main

import ("fmt"
"math")

func main(){

	test := map[string]int{
		"kapil":2,
		"raj":3,
		"rajni" :5,
		"kulbhushan":6,
	}
	//test["raj"] = 0
	delete(test,"raj")
	for key,value := range test{
		fmt.Println(key,value)
	}
	a := math.MaxInt32
	fmt.Println(a)
}