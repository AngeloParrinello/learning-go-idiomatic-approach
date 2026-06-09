package main

import (
	"fmt"
	"time"
)

// We can also use the done channel pattern to implement a pattern that we previously saw:
// return a cancellation function alongside the channel.

func main() {
	ch, cancel := countTo(10)
	for i := range ch {
		if i > 5 {
			break
		}
	}
	cancel()

	// Blocking a bit the main goroutine, we can see the "done" print
	time.Sleep(1 * time.Second)
}

// The countTo function creates two channels, one that returns data and another for signaling
// done. Rather than return the done channel directly, we create a closure that
// closes the done channel and return the closure instead. Cancelling with a closure
// allows us to perform additional clean-up work, if needed.
func countTo(max int) (<-chan int, func()) {
	ch := make(chan int)
	done := make(chan struct{})
	cancel := func() {
		close(done)
	}

	go func() {
		for i := 0; i < max; i++ {
			select {
			case <-done:
				fmt.Println("done")
				// This return statement is necessary to terminate the goroutine when the done channel is closed.
				// If we don't return here, the for-loop will continue and at every iteration will print "done"
				// until it reaches the max value and closes the channel, which is not what we want.
				// es i = 10
				// sent 0
				// sent 1
				// sent 2
				// sent 3
				// sent 4
				// sent 5
				// sent 6
				// done
				// done
				// done
				return
			case ch <- i:
				fmt.Println("sent", i)
			}
		}
		close(ch)
	}()

	return ch, cancel
}
