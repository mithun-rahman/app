package channel

import (
	"fmt"
	"time"
)

func doJob(done <-chan interface{}) <-chan interface{} {
	terminated := make(chan interface{})

	go func() {
		defer fmt.Println("doJob exited")
		defer close(terminated)

		for {
			select {
			case <-done:
				return
			default:
				fmt.Println("doJob running")
			}

		}
	}()

	fmt.Println("doJob initiate")

	return terminated
}

func Terminated() {
	done := make(chan interface{})
	terminated := doJob(done)

	go func() {
		// Cancel the operation after 1 second.
		time.Sleep(5 * time.Second)
		fmt.Println("Canceling doWork goroutine...")
		// if it has no cancel produces a deadlock
		close(done)
	}()

	fmt.Println("initiate ...")
	d := <-terminated
	fmt.Println(d)
	fmt.Println("Done.")

}
