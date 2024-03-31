package letsgo

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
)

type HashReader interface {
	io.Reader
	hash() string
}

func hashAndBroadcast(r *hashReader) error {
	hash := r.hash()
	fmt.Println(hash)
	return broadCast(r)
}

func broadCast(r io.Reader) error {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("string of bytes:", string(b))
	return nil
}

type hashReader struct {
	*bytes.Reader
	buf *bytes.Buffer
}

func NewHashReader(b []byte) *hashReader {
	return &hashReader{
		Reader: bytes.NewReader(b),
		buf:    bytes.NewBuffer(b),
	}
}

func (h *hashReader) hash() string {
	return hex.EncodeToString(h.buf.Bytes())
}

func Interface() {
	payload := []byte("hello world")
	hashAndBroadcast(NewHashReader(payload))
}
