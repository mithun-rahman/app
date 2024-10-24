package concurrency

import "fmt"

var count int

func incCount() {
	count++
}

func doCount() {
	for i := 0; i < 1000; i++ {
		go incCount()
	}
}

func Pessimistic() {
	count = 0
	doCount()

	fmt.Println(count)
}
