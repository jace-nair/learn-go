package main

import (
	"fmt"
	"sync"
)

func worker1(c chan<- string, wg *sync.WaitGroup) {
	//wg.Add(1)
	defer wg.Done()
	sum := 0
	for i := 0; i < 100; i++ {
		//fmt.Println("IDX from worker1:", i)
		sum += i
	}
	c <- fmt.Sprintf("This is the sum from worker1: %v", sum)
}

func worker2(c chan<- string, wg *sync.WaitGroup) {
	//wg.Add(1)
	defer wg.Done()
	sum := 0
	for i := 0; i < 100; i++ {
		//fmt.Println("IDX from worker2:", i)
		sum += i
	}
	c <- fmt.Sprintf("This is the sum from worker2: %v", sum)
}
