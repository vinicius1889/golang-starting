package main

import (
	"fmt"
	"time"
	"runtime"
	"sync"
)

func init(){ runtime.GOMAXPROCS(runtime.NumCPU())}

var wg sync.WaitGroup

var mystart2 int64

func main(){

	mystart2 = time.Now().UnixNano()
	wg.Add(2)
	go concur3()
	go concur4()
	wg.Wait()
}

func concur3(){
	for i:=0;i<100000;i++ {
		//time.Sleep(200)
	}
	fmt.Println("concur3 ",time.Now().UnixNano()-mystart2)
	wg.Done()
}

func concur4(){
	for i:=0;i<100000;i++ {
		//fmt.Print("Concur 2 line = ",i)
	}
	fmt.Println("concur4 ",time.Now().UnixNano()-mystart2)
	wg.Done()
}

