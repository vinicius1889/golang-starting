package main

import "fmt"

func setToZeroUsingPointer(variable *int){
	*variable = 0
}

func setToZero(variable int){
	variable=0
	fmt.Println(&variable)
}

func main(){
	num1 := 10
	setToZeroUsingPointer(&num1)
	fmt.Println(num1)

	num2:= 20
	setToZero(num2)
	fmt.Println(&num2)
	fmt.Println(num2)
}
