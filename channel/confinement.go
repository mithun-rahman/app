package channel

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

func producer(ch chan int) {
	fmt.Println("producer start")
	defer close(ch)
	for i := 0; i < 10; i++ {
		ch <- i
	}
}

func consumer(ch chan int) {
	for v := range ch {
		fmt.Println(v)
	}
	fmt.Println("consumer done")
}

func AdhocConfinement() {
	ch := make(chan int)
	go producer(ch)
	go consumer(ch)
}

func addData() <-chan int {
	results := make(chan int, 5)

	go func() {
		defer close(results)

		for i := 0; i <= 5; i++ {
			results <- i
		}
	}()

	return results
}

func getData(results <-chan int) {
	for v := range results {
		fmt.Println(v)
	}
}

func LexicalConfinement() {
	ch := addData()
	go getData(ch)
}

type MtbCall func(d *Data, ctx context.Context, wg *sync.WaitGroup, errChan chan error, cancelFunc context.CancelFunc)

type Data struct {
	Nom int
	Per int
	Doc string
}

func AddNominee(d *Data, ctx context.Context, wg *sync.WaitGroup, errChan chan error, cancelFunc context.CancelFunc) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("AddNominee: error occurred")
			return
		default:
			time.Sleep(2 * time.Second)
			fmt.Println("AddNominee: ", d.Nom)
		}
	}
}

func AddData(d *Data, ctx context.Context, wg *sync.WaitGroup, errChan chan error, cancelFunc context.CancelFunc) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("AddData : error occurred")
			return
		default:
			time.Sleep(2 * time.Second)
			fmt.Println("AddData: ", d.Per)
		}
	}
}

func UploadDocument(d *Data, ctx context.Context, wg *sync.WaitGroup, errChan chan error, cancelFunc context.CancelFunc) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("UploadDocument : error occurred")
			return
		default:
			fmt.Println("UploadDocument: ", d.Doc)
			errChan <- errors.New("UploadDocument error")
			cancelFunc()
		}
	}
}

func (d *Data) Nominee(body int) {
	d.Nom = body
}

func (d *Data) Document(body string) {
	d.Doc = body
}

func (d *Data) Person(body int) {
	d.Per = body
}

func Main() {
	var wg sync.WaitGroup
	errChan := make(chan error)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	dt := &Data{}
	dt.Nominee(5)
	dt.Document("mithun")
	dt.Person(89)

	mtbCall := []MtbCall{AddData, AddNominee, UploadDocument}

	for _, mtb := range mtbCall {
		wg.Add(1)
		go mtb(dt, ctx, &wg, errChan, cancel)
	}

	go func() {
		defer close(errChan)
		wg.Wait()
	}()

	for err := range errChan {
		if err != nil {
			fmt.Println(err)
		}
	}

	fmt.Println("All goroutines have completed.")

}

func RetryFunc(ret int) error {
	if ret == 0 {
		return nil
	}
	err := errors.New("retry function failed")
	if err == nil {
		return err
	}
	return RetryFunc(ret - 1)
}
