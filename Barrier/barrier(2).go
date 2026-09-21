//Barrier.go Template Code
//Copyright (C) 2024 Dr. Joseph Kehoe

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

//--------------------------------------------
// Author: Joseph Kehoe (Joseph.Kehoe@setu.ie)
// Created on 30/9/2024
// Modified by: Dorian Nowacki (C00288613)
// Issues: ZERO!!!!
// The barrier is not implemented!
//--------------------------------------------

//CLASSMATES I WORKED WITH:
//Mark lambert (brains of the operation)
//Amelia Hamulewicz (judging me >:( )
//Adam Noonan (Adam is a cool guy)

// CLASSMATES I HELPED:
// Sean McGrath (Passing on what i learnt)
// Stella
package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/semaphore"
)

// ///////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Added these Vars
var countlock sync.Mutex             //Mutex - Lets you put the lock on the counter variable
var barrier = make(chan struct{}, 1) //Semaphore - Used to block every thread until the last is ready
var count = 0                        // track how many threads there are
/////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// Place a barrier in this function --use Mutex's and Semaphores
func doStuff(goNum int, wg *sync.WaitGroup) bool {
	time.Sleep(time.Second)
	fmt.Println("Part A", goNum)

	countlock.Lock() // Locking the count variable until reading or changing
	count++          // innit

	if count == 10 { //still locked - checking if count variable reached 10
		barrier <- struct{}{} // when count is 10 - Last thread signals to the other for enterance
	}

	countlock.Unlock()    //Finished so unlocking mutex
	<-barrier             //wait
	barrier <- struct{}{} //signal

	//we wait here until everyone has completed part A
	fmt.Println("PartB", goNum)
	wg.Done()
	return true
}

// main stays the same
func main() {
	totalRoutines := 10
	var wg sync.WaitGroup
	wg.Add(totalRoutines)
	//we will need some of these
	ctx := context.TODO()
	var theLock sync.Mutex
	sem := semaphore.NewWeighted(int64(totalRoutines))
	theLock.Lock()
	sem.Acquire(ctx, 1)
	for i := range totalRoutines { //create the go Routines here
		go doStuff(i, &wg)
	}
	sem.Release(1)
	theLock.Unlock()

	wg.Wait() //wait for everyone to finish before exiting
}

/*without code modification:
Output was: ABABABABAB

with code modification:
Output is: AAAAABBBBB
*/
