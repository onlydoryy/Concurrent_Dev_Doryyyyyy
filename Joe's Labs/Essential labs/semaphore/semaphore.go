//////////////////////////////////////////////////////////////////////////////////
//Dorian Nowacki
//C00288613
//21-09-2026
//////////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"sync"
	"time"
)

// make struct containing channel
// add init, acquire and release
type semaphore struct {
	theCounter chan struct{}
}

//space added at func and "Acqite" changed to "Acquire" -- "Semaphore" changed to "semaphore" -- //!!!Amelia Hamulewicz!!! spotted there was no curly brackets
/*func Acquire(sem *semaphore) {

}

*/
func main() {
	maxGoroutines := 5
	semaphore := make(chan struct{}, maxGoroutines)

	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			semaphore <- struct{}{}
			defer func() { <-semaphore }()

			// Simulate a task
			fmt.Printf("Running task %d\n", i)
			time.Sleep(2 * time.Second)
		}(i)
	}
	wg.Wait()
}

/* Output with unused function commented out:
Running task 0
Running task 2
Running task 1
Running task 10
Running task 19
Running task 3
Running task 4
Running task 5
Running task 6
Running task 7
Running task 8
Running task 9
Running task 14
Running task 15
Running task 11
Running task 12
Running task 16
Running task 17
Running task 13
Running task 18
*/
