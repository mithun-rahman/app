package letsgo

import (
	"fmt"
	"net/http"
)

func myExecuteFunc(db DB) ExecuteFn {
	return func(s string) {
		fmt.Println("execute my func:", s)
		db.Store(s)
	}
}

type DB interface {
	Store(string) error
}

type Store struct{}

func (s *Store) Store(value string) error {
	fmt.Println("storing in db: ", value)
	return nil
}

func makeHTTPFunc(db DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		db.Store("hello from handler")
	}
}

type ExecuteFn func(string)

func Execute(fn ExecuteFn) {
	fn("hello world")
}
