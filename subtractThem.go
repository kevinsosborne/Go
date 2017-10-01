package main

import "fmt"

func main(){
	fmt.Println(substractThen(1,2,3,4,5))
}

func substractThem(args ...int) int{
	finalValue := 0
	for _, value := range args {
		finalValue -= value
	}
	return finalValue
}