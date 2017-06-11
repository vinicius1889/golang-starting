package main

import (
	"sync"
	"fmt"
	"time"
)

var waitGroup sync.WaitGroup

var mystart int64

func main(){
	mystart = time.Now().UnixNano()
	waitGroup.Add(2)
	go concur1()
	go concur2()
	waitGroup.Wait()

}

func concur1(){
	for i:=0;i<100000;i++ {
	//	fmt.Println("Concur 1 line = ",i)
		//time.Sleep(200)
	}
	fmt.Println("Concur1 ",time.Now().UnixNano()-mystart)
	waitGroup.Done()
}

func concur2(){
	for i:=0;i<100000;i++ {
	//	fmt.Println("Concur 2 line = ",i)
	}
	fmt.Println("Concur2 ",time.Now().UnixNano()-mystart)
	waitGroup.Done()
}

