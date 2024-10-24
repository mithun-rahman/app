package channel

import (
	"fmt"
	"sync"
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

// type MtbCall func(d *Data, wg *sync.WaitGroup, errChan chan error, done chan bool)
//
//	type Data struct {
//		Nom int
//		Per int
//		Doc string
//	}
//
//	func AddNominee(d *Data, wg *sync.WaitGroup, errChan chan error, done chan bool) {
//		defer wg.Done()
//		for {
//			select {
//			case <-done:
//				fmt.Println("AddNominee: error occurred")
//				return
//			default:
//				time.Sleep(2 * time.Second)
//				fmt.Println("AddNominee: ", d.Nom)
//			}
//		}
//	}
//
//	func AddData(d *Data, wg *sync.WaitGroup, errChan chan error, done chan bool) {
//		defer wg.Done()
//		for {
//			select {
//			case <-done:
//				fmt.Println("AddData : error occurred")
//				return
//			default:
//				fmt.Println("AddData: ", d.Per)
//				time.Sleep(5 * time.Second)
//				fmt.Println("AddData1: ", d.Per)
//			}
//		}
//	}
//
//	func UploadDocument(d *Data, wg *sync.WaitGroup, errChan chan error, done chan bool) {
//		defer wg.Done()
//		for {
//			select {
//			case <-done:
//				fmt.Println("UploadDocument : error occurred")
//				return
//			default:
//				fmt.Println("UploadDocument: ", d.Doc)
//				errChan <- errors.New("UploadDocument error")
//				close(done)
//			}
//		}
//	}
//
// func main() {
var wg sync.WaitGroup

//	errChan := make(chan error)
//	done := make(chan bool)
//
//	dt := &Data{}
//	dt.Nom = 5
//	dt.Doc = "mithun"
//	dt.Per = 9
//
//	mtbCall := []MtbCall{AddData, AddNominee, UploadDocument}
//
//	for _, mtb := range mtbCall {
//		wg.Add(1)
//		go mtb(dt, &wg, errChan, done)
//	}
//
//	go func() {
//		defer close(errChan)
//		wg.Wait()
//	}()
//
//	for err := range errChan {
//		if err != nil {
//			fmt.Println(err)
//		}
//	}
//
//	fmt.Println("All goroutines have completed.")
//
//}
//
//func RetryFunc(ret int) error {
//	if ret == 0 {
//		return nil
//	}
//	err := errors.New("retry function failed")
//	if err == nil {
//		return err
//	}
//	return RetryFunc(ret - 1)
//}
