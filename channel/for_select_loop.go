package channel

import "fmt"

func ForSelectLoop() {
	charChannel := make(chan string, 3)
	chars := []string{"a", "b", "c"}

	for _, s := range chars {
		select {
		case charChannel <- s:

		}
	}

	close(charChannel)

	for _, s := range chars {
		fmt.Println(s)
	}
}
