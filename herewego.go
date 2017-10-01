package main

import "fmt"

// Comments

func main(){


	// var (
	// 	var A = 2;
	// 	var B = 3;
	// )

	
	// STRINGS
	// var myName string = "Kevin";
	// fmt.Println(len(myName));
	// fmt.Println(myName + " is a person \nnew line");

	// BOOLEAN
	// const pi float64 = 3.14159265;	
	// var isOver bool = true;
	// shows only the first three decimial places
	// fmt.Printf("%.3f \n", pi);
	// t is the value
	// fmt.Printf("%t \n", isOver);
	// T is the type
	// fmt.Printf("%T \n", isOver);


	// different ways to represent a value
	// %d is the digit
	// fmt.Printf("%d \n", 100);
	// %b is the binary code
	// fmt.Printf("%b \n", 100);
	// %c is the number of digits
	// fmt.Printf("%c \n", 100);
	// %x is the hex code
	// fmt.Printf("%x \n", 100);
	// % is the scientic code
	// fmt.Printf("%e \n", 100);

	// true and false statements
	// fmt.Println("true && false = ", true && false);
	// fmt.Println("true || false = ", true || false);
	// fmt.Println("!true = ", !true);

	// For Loops
	// i := 1
	// for i <= 10 {
	// 	fmt.Println(i);
	// 	i++;
	// }

	// for j := 0; j < 5; j++{
	// 	fmt.Println(j);
	// }


	// If Statements
	// yourAge := 18
	// if yourAge >= 16 && yourAge < 18 {
	// 	fmt.Println("You Can Drive")
	// } else if yourAge >= 18{
	// 	fmt.Println("You Can Vote")
	// } else {
	// 	fmt.Println("You can have fun");
	// }

	// Switch Statements
	// yourAge := 18	
	// switch yourAge {
	// case 16 : fmt.Println("Go Drive")
	// case 18 : fmt.Println("Go Vote")
	// default : fmt.Println("Go Have Fun")
	// }

	// var favNums2[5] float64
	// favNums2[0] = 163
	// favNums2[1] = 785655;
	// favNums2[2] = 691
	// favNums2[3] = 3.141
	// favNums2[4] = 1.618;

	// fmt.Println(favNums2[3])
	
	// favNums3 := [5]float64 {23,2,34,4,5}
	// for _, value := range favNums3 {
	// 	fmt.Println(value)
	// }
	// for i, value := range favNums3 {
	// 	fmt.Println(i, value)
	// }

	// Slice an Array
	// numSlice := []int {5,4,3,2,1}
	// numSlice2 := numSlice[3:5]
	// fmt.Println("numSlice2[0] = ", numSlice2[0])
	// for _, value := range numSlice2 {
	// 	fmt.Println(value)
	// }
	// fmt.Println("numSlice[3:5] = ", numSlice[3:5]);
	// numSlice := []int {0,1,2,3,4,5}
	// make(creates a new array with number of element spaces, and then additional number of elements you can add on)
	// numSlice3 := make([]int, 6,6)
	// copy(numSlice3, numSlice)
	// for _, value := range numSlice3 {
	// 	fmt.Println(value)
	// }
	// numSlice3 = append(numSlice3, 3);
	// for _, value := range numSlice3 {
	// 	fmt.Println(value)
	// }
	// fmt.Println(numSlice3[5])

	// MAPS
	// presAge := make(map[string] int)
	// presAge["TheordoreRoosevelt"] = 42
	// fmt.Println(len(presAge))
	// presAge["John F. Kennedy"] = 43
	// fmt.Println(len(presAge))
	// delete(presAge, "John F. Kennedy")
	// fmt.Println(len(presAge))
	
	// NUMBERS
	listNums := []float64{1,2,3,4,5}

	// fmt.Println("Hello World");

	// var age int = 40;
	// var favNum float64 = 1.6180339;

	// fmt.Println(age, " ", favNum);

	// var numOne = 1.000;
	// var num99 = .9999

	// fmt.Println(numOne - num99);


	// fmt.Println("6 + 4 =", 6 + 4);
	// fmt.Println("6 * 4 =", 6 * 4);
	// fmt.Println("6 - 4 =", 6 - 4);
	// fmt.Println("6 / 4 =", 6 / 4);
	// fmt.Println("6 % 4 =", 6 % 4);
}