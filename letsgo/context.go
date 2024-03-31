package letsgo

import (
	"context"
	"fmt"
	"time"
)

func CallContext() {
	start := time.Now()

	ctx := context.Background()
	userID := 10
	id, err := fetchUserData(ctx, userID)

	fmt.Println(id)
	fmt.Println(err)

	duration := time.Since(start)
	fmt.Println(duration)
}

type Response struct {
	value int
	err   error
}

func fetchUserData(ctx context.Context, userID int) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*700)
	defer cancel()

	respChan := make(chan Response)
	go func() {
		val, err := ThirdPartyStuff()
		respChan <- Response{
			value: val,
			err:   err,
		}
	}()
	for {
		select {
		case <-ctx.Done():
			return -1, fmt.Errorf("third party timed out")
		case resp := <-respChan:
			return resp.value, resp.err
		}
	}
}

func ThirdPartyStuff() (int, error) {
	time.Sleep(time.Millisecond * 750)
	return 170225, nil
}
