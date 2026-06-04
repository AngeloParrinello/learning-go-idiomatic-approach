package example_chapter_10

import "fmt"

func main() {
	// Concurrency in Go is based on CSP (Communicating Sequential Processes)

	// A process is an instance of a program that’s being run by a computer’s operating system.
	// A process is composed of one or more threads.
	// Threads within a process share access to resources. A CPU can execute instructions
	// from one or more threads at the same time, depending on the number of cores. One
	// of the jobs of an operating system is to schedule threads on the CPU to make sure
	// that every process (and every thread within a process) gets a chance to run.

	// Goroutines are lightweight processes managed by the Go runtime. When a Go program starts,
	// the Go runtime creates a number of threads and launches a single
	// goroutine to run your program. All of the goroutines created by your program,
	// including the initial one, are assigned to these threads automatically by the Go runtime scheduler,
	// just as the operating system schedules threads across CPU cores.

	// Any function can be launched as a goroutine.
	// it is customary in Go to launch goroutines with a closure
	// that wraps business logic. The closure takes care of the concurrent bookkeeping.

	// Goroutines communicate using channels. Like slices and maps, channels are a built-in
	// type created using the make function:
	ch := make(chan int)
	// Like maps, channels are reference types. When you pass a channel to a function, you
	// are really passing a pointer to the channel. Also like maps and slices, the zero value
	// for a channel is nil.

	// Use the <- operator to interact with a channel.
	a := <-ch // reads a value from ch and assigns it to a
	fmt.Println(a)
	var b int
	ch <- b // write the value in b to ch

	// Each value written to a channel can only be read once. If multiple goroutines are
	// reading from the same channel, a value written to the channel will only be read by
	// one of them.

	// It is rare for a goroutine to read and write to the same channel
	// use an arrow before the chan keyword (ch <-chan int) to indicate that the goroutine only reads from the channel
	// use an arrow after the chan keyword (chan<- int) to indicate that the goroutine only writes to the channel

	// By default channels are unbuffered. Every write to an open, unbuffered channel
	//causes the writing goroutine to pause until another goroutine reads from the same
	//channel. Likewise, a read from an open, unbuffered channel causes the reading
	// goroutine to pause until another goroutine writes to the same channel. This means
	// you cannot write to or read from an unbuffered channel without at least two concurrently
	// running goroutines.

	// Go also has buffered channels. These channels buffer a limited number of writes
	// without blocking. Just as writing to a channel with a full buffer blocks,
	// reading from a channel with an empty buffer also blocks.
	ch = make(chan int, 10)

	// Most of the time, you should use unbuffered channels.

	// Read from a channel using a for-range loop
	for v := range ch {
		// The loop continues until the channel is closed, or until a break or
		// return statement is reached.
		fmt.Println(v)
	}

	// When you are done writing to a channel, you close it using the built-in close
	// function:
	close(ch)

	// how do we tell the difference between a zero value that
	// was written and a zero value that was returned because the channel is closed?
	v, ok := <-ch
	// If ok is true -> channel open, otherwise closed
	fmt.Println(v, ok)

	// the standard pattern is to make the writing goroutine responsible for closing the channel
	// when there’s nothing left to write.
	// When multiple goroutines are writing to the same
	// channel, this becomes more complicated, as calling close twice on the same channel
	// causes a panic. Furthermore, if you close a channel in one goroutine, a write to the
	// channel in another goroutine triggers a panic as well. The way to address this is to use
	// a sync.WaitGroup.

	// select statement is the control structure for concurrency in Go, and solves the question:
	// if you can perform two concurrent operations, which one do you do first?
	// The select keyword allows a goroutine to read from or write to one of a set of multiple channels.

	// Another advantage of select choosing at random is that it prevents one of the most
	// common causes of deadlocks: acquiring locks in an inconsistent order. If you have
	// two goroutines that both access the same two channels, they must be accessed in the
	// same order in both goroutines, or they will deadlock.Another advantage of select choosing at random is that it prevents one of the most
	// common causes of deadlocks: acquiring locks in an inconsistent order. If you have
	// two goroutines that both access the same two channels, they must be accessed in the
	// same order in both goroutines, or they will deadlock.

	// Remember that our main is running on a goroutine that is launched at startup by the
	// Go runtime.

	// Since select is responsible for communicating over a number of channels, it is often
	// embedded within a for loop:
	//for {
	//	select {
	//	case <-done:
	//		return
	//	case v := <-ch:
	//		fmt.Println(v)
	//	}
	//}
	// This is so common that the combination is often referred to as a for-select loop.
	// When using a for-select loop, you must include a way to exit the loop.

	// If you want to implement a nonblocking read or write on a channel,
	// use a select with a default. The following code does not wait if there’s no value to
	// read in ch; it immediately executes the body of the default:
	select {
	case v := <-ch:
		fmt.Println(v)
	default:
		fmt.Println("no value to read")
	}
	// Having a default case inside a for-select loop is almost always
	// the wrong thing to do. It will be triggered every time through the
	// loop when there’s nothing to read or write for any of the cases. This
	// makes your for loop run constantly, which uses a great deal of CPU.

	// Keep Your APIs Concurrency-Free
	// this means that you should never expose channels or mutexes in your
	// API’s types, functions, and methods

	// Goroutines, for Loops, and Varying Variables
	// Most of the time, the closure that you use to launch a goroutine has no parameters.
	// There is one common situation where this doesn’t work: when trying to capture the index or value
	// of a for loop.  This code contains a subtle bug:
	x := []int{2, 4, 6, 8, 10}
	ch3 := make(chan int, len(x))
	for _, v := range x {
		go func() {
			ch3 <- v * 2
		}()
	}
	for i := 0; i < len(x); i++ {
		fmt.Println(<-ch3)
	}
	// We launch one goroutine for each value in a. It looks like we pass a different value in
	// to each goroutine, but running the code shows something different:
	// 20
	// 20
	// 20
	// 20
	// 20
	// The reason why every goroutine wrote 20 to ch is that the closure for every goroutine captured the same variable.
	// Possible solution is to pass the variable to the closure:
	/* for _, v := range a {
		go func(val int) {
			ch <- val * 2
		}(v)
	}*/

	// Always Clean Up Your Goroutines
	// Whenever you launch a goroutine function, you must make sure that it will eventually exit.
	// Unlike variables, the Go runtime can’t detect that a goroutine will never be used again.
	// If a goroutine doesn’t exit, the scheduler will still periodically give it time to do nothing,
	// which slows down your program. This is called a goroutine leak.
}
