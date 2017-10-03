package main

import "fmt"

func main(){
	defer printTwo()
	printOne()
	printThree()
}

func printOne(){ fmt.Println(1)}
func printTwo(){ fmt.Println(2)}
func printThree(){ fmt.Println(3)}