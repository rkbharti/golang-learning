package main

import (
	"fmt"
	"sync"
	"time"
)

func firstHello(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Hello World from first function")
	time.Sleep(time.Second * 3)
}
func secondHello(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Hello World from second function")
	time.Sleep(time.Second * 3)
}
func taskCompleted(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <= 10; i++ {
		fmt.Println("task done is :", i)
		time.Sleep(time.Second * 3)
	}
}

func main() {
	var wg sync.WaitGroup
	fmt.Println("Learing Go-Routines")
	wg.Add(3)
	go firstHello(&wg)
	go secondHello(&wg)
	go taskCompleted(&wg)

	wg.Wait()

}
