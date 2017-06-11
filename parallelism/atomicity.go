package main

import (
	"fmt"
	"time"
	"sync"
	"sync/atomic"
)

var countConcurrency int32 = 0

var wg2 sync.WaitGroup

func main(){
	wg2.Add(3)
	go concurrencyTeste("Thread1 ")
	go concurrencyTeste("Thread2 ")
	go concurrencyTeste("Thread3 ")
	wg2.Wait()
}

func concurrencyTeste(threadName string){
	for i:=0;i<100;i++ {
		atomic.AddInt32( &countConcurrency , 1)
		time.Sleep(300)
		fmt.Println("Nome = ",threadName," count = ",countConcurrency)
	}
	wg2.Done()
}

