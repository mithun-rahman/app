package channel

import "fmt"

func SliceToChannel(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, num := range nums {
			out <- num
		}
		close(out)
	}()
	return out
}

func NumToSquare(nums <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for num := range nums {
			out <- num * num
		}
		close(out)
	}()
	return out
}

func PipelineChannel() {
	nums := []int{2, 3, 1, 8, 4}
	dataChannel := SliceToChannel(nums)
	//for num := range dataChannel {
	//	fmt.Println(num)
	//}
	finalChannel := NumToSquare(dataChannel)

	for num := range finalChannel {
		fmt.Println(num)
	}

	for num := range finalChannel {
		fmt.Println(num)
	}
}
