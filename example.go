package main

import (
	"github.com/matthieuEv/go-sha256/sha256"
	"fmt"
	"io"
	"os"
)

func example(){
	// sha256.DEBUG = true
	var msg = []byte("The quick brown fox jumps over the lazy dog")
	sha := sha256.Sum256(msg)
	fmt.Println(sha)
	var msg2 = []byte("hello world")
	sha2 := sha256.Sum256(msg2)
	fmt.Println(sha2)

	// Test with a file
	f, err := os.Open("sha256/const.go")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}
	fmt.Println(sha256.Sum256(data))
}