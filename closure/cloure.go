package main

import "fmt"

func clTest(a int , b int) int { return a+b }

func clClosure() func(a int , b int ) int{
	return func(num1 int , num2 int) int {
		return num1+num2
	}
}

func clClosure2(num1 int, num2 int) int {
    var result int
	{

		resultCl := num1+num2
		result = resultCl
	}
	//fmt.Println(resultCl)
	return result
}

func main(){
	soma := clClosure()
	result := soma(10,20)


	fmt.Println(clTest(1,2))
	fmt.Println(result)
	fmt.Println(clClosure2(100,200))

}
