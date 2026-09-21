//CLASSMATES I WORKED WITH:
//Mark lambert (brains of the operation)
//Amelia Hamulewicz (judging me >:( )
//Adam Noonan (Adam is a cool guy)

// CLASSMATES I HELPED:
// Sean McGrath (Passing on what i learnt)
// Stella

package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

// Global variables shared between functions --A BAD IDEA
// Moved these to make them Gobal (MR WORLDWIDE)
var aArrived = make(chan struct{})
var bArrived = make(chan struct{})

// Joeseph code
func WorkWithRendezvous(wg *sync.WaitGroup, Num int) bool {
	var X time.Duration
	X = time.Duration(rand.IntN(5))

	time.Sleep(X * time.Second) //wait random time amount
	fmt.Println("Part A", Num)
	//Rendezvous here
	aArrived <- struct{}{} //A sends a signal to B

	<-bArrived // B waits for A
	fmt.Println("PartB", Num)
	wg.Done()
	return true
}

func main() {
	//barrier := make(chan bool)
	var wg sync.WaitGroup
	threadCount := 5

	wg.Add(threadCount)
	for N := range threadCount {

		go WorkWithRendezvous(&wg, N)
	}

	//Added 2 for loops
	for range threadCount {
		<-aArrived //A's are waiting for B
	}
	for range threadCount {
		<-bArrived //B's are sending a signal to all waiting A's
	}

	wg.Wait() //wait here until everyone (10 go routines) is done

}

/*without code modification:
Output was: ABABABABAB

with code modification:
Output is: AAAAABBBBB
*/
