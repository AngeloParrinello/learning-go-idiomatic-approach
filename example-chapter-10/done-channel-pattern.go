package main

import (
	"fmt"
	"time"
)

func main() {
	searchGoogle := func(s string) []string {
		fmt.Println(s)
		time.Sleep(1 * time.Second)

		return []string{"google result 1", "google result 2", "google result 3"}
	}

	searchBing := func(s string) []string {
		fmt.Println(s)
		time.Sleep(2 * time.Second)

		return []string{"bing result 1", "bing result 2", "bing result 3"}
	}

	searchDuckDuckGo := func(s string) []string {
		fmt.Println(s)
		time.Sleep(3 * time.Second)

		return []string{"duck result 1", "duck result 2", "duck result 3"}
	}

	searchers := []func(string) []string{
		searchGoogle,
		searchBing,
		searchDuckDuckGo,
	}

	result := searchData2("test", searchers)

	fmt.Println(result)

}

func searchData2(s string, searchers []func(string) []string) []string {
	// we declare a channel named done that contains data of type
	// struct{}. We use an empty struct for the type because the value is unimportant; we
	// never write to this channel, only close it.
	done := make(chan struct{})
	result := make(chan []string)
	for _, searcher := range searchers {
		// We launch a goroutine for each searcher
		// passed in. The select statements in the worker goroutines wait for either a write on
		// the result channel (when the searcher function returns) or a read on the done
		// channel.
		go func(searcher func(string) []string) {
			select {
			// Remember that a read on an open channel pauses until there is data available
			// and that a read on a closed channel always returns the zero value for the channel.
			// This means that the case that reads from done will stay paused until done is closed. In
			// searchData, we read the first value written to result, and then we close done. This
			// signals to the goroutines that they should exit, preventing them from leaking.
			case result <- searcher(s):
				fmt.Println("searcher returned result")
			case <-done:
				fmt.Println("searcher done")
			}
		}(searcher)
	}

	// We return the first result immediately and then close the done channel to signal to the other goroutines that they should exit.
	// This is an example of the done channel pattern, which is a common way to signal to goroutines that they should stop doing work and exit.
	r := <-result
	close(done)

	return r
}
