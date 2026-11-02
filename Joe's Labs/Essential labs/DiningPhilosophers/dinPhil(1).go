
// Dining Philosophers Problem
// Author: Joseph Kehoe
// Created: 21/10/24
// GPL Licence
//
// Dorian Nowacki
// C00288613
// 02-11-2026

package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func think(index int) {
	var X time.Duration
	X = time.Duration(rand.IntN(5))
	time.Sleep(X * time.Second)
	fmt.Println("Phil:", index, "was thinking")
}

func eat(index int) {
	var X time.Duration
	X = time.Duration(rand.IntN(5))
	time.Sleep(X * time.Second)
	fmt.Println("Phil:", index, "was eating")
}

func getForks(index int, forks map[int]chan bool) {
	if index == 4 {
		forks[(index+1)%5] <- true
		forks[index] <- true
	} else {
		forks[index] <- true
		forks[(index+1)%5] <- true
	}
}

func putForks(index int, forks map[int]chan bool) {
	<-forks[index]
	<-forks[(index+1)%5]
}

func doPhilStuff(index int, wg *sync.WaitGroup, forks map[int]chan bool) {
	defer wg.Done()

	for i := 0; i < 5; i++ {
		think(index)
		getForks(index, forks)
		eat(index)
		putForks(index, forks)
	}
}

func main() {
	var wg sync.WaitGroup

	philCount := 5
	wg.Add(philCount)

	forks := make(map[int]chan bool)

	for k := range philCount {
		forks[k] = make(chan bool, 1)
	} // Set up forks

	for N := range philCount {
		go doPhilStuff(N, &wg, forks)
	} // Start philosophers

	wg.Wait() // Wait until all five goroutines finish

	fmt.Println("All philosophers have finished eating.")
}
