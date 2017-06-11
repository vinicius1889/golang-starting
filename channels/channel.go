package main

import (
	"fmt"
	"sync"
)

func main(){

	channel := make(chan int)
	var wgChannels sync.WaitGroup
	wgChannels.Add(2)

	go func(){
		for i:= 0;i<10;i++{
			channel<-i
		}
		wgChannels.Done()
	}()

	go func(){
		for i:= 10;i<20;i++{
			channel<-i
		}
		wgChannels.Done()
	}()

	go func(){
		wgChannels.Wait()
		close(channel)
	}()

	for n := range channel{
		fmt.Println(n)
	}



}
