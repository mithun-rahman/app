package channel

import (
	"fmt"
	"time"
)

func doWork(done <-chan bool) {
	for {
		select {
		case <-done:
			fmt.Println("Done")
			return
		default:
			time.Sleep(time.Millisecond * 500)
		}
	}
}

func DoneChannel() {
	done := make(chan bool)

	go doWork(done)

	time.Sleep(time.Second * 2)

	close(done)
}
