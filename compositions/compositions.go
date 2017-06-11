package main

import "fmt"


type People struct{
	name string
	age int
}

func (p People) sayHello() string{
	return "Scarface: Say Hello to my little friend ---> "+p.name
}

func main(){
	people := People{ 	name:"vinicius dias", age:28	}
	fmt.Println(people.sayHello())
}
