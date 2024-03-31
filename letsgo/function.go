package letsgo

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type TransformFunc func(string) string

type Server struct {
	FilenameTransformFunc TransformFunc
}

func (s *Server) HandleRequest(filename string) error {
	newFilename := s.FilenameTransformFunc(filename)
	fmt.Println("newFilename: ", newFilename)
	return nil
}

func HashFilename(filename string) string {
	hash := sha256.Sum256([]byte(filename))
	newFilename := hex.EncodeToString(hash[:])
	return newFilename
}

func PrefixFilename(prefix string) TransformFunc {
	return func(filename string) string {
		return prefix + filename
	}
}
