package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup
var counter int
var mutex sync.Mutex

func main(){
	wg.Add(2)
	go increamentor("Foo: ")
	go increamentor("Bar: ")
	wg.Wait()
	fmt.Println("Final counter : ", counter)
}

func increamentor(s string){
	for i:=0; i<20; i++{
		time.Sleep(time.Duration(rand.Intn(20))*time.Millisecond)
		mutex.Lock()
		counter++
		fmt.Println(s, i, " Counter : ", counter)
		mutex.Unlock()
	}
	wg.Done()
}
