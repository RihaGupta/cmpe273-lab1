package main

import "fmt"

func main() {    
	var n int16
	fmt.Print("Enter number:")
	fmt.Scanf("%d",&n)
	fmt.Print("number is ")
	fmt.Println(n)
	fmt.Println(fibonacci(n))
}

func fibonacci(n int16) int{
	if n<0 {
		return -1
	} else if n==0 {
		return 0
	} else if n==1 {
		return 1
	} else {
		return fibonacci(n-1)+fibonacci(n-2)
	}
}