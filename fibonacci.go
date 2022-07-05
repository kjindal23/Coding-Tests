package main

import "fmt"

func fibonacci(num int)int{

	if num <= 0{
		return 0
	}

	slNum := make([]int,num+1)

	slNum[0]=0
	slNum[1]=1
	for i := 2; i<=num; i++ {
		slNum[i] = slNum[i-1] + slNum[i-2]
		fmt.Println(slNum[i])
	}
	return slNum[num]
}

func main(){

	x := fibonacci(-2)

   fmt.Println("fib number = " , x)
}