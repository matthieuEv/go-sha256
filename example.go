package main

import "github.com/matthieuEv/go-sha256/sha256"

func example(){
	// sha256.DEBUG = true
	var msg = []byte("The quick brown fox jumps over the lazy dog")
	sha := sha256.Sum256(msg)
	print(sha)
	print("\n")
	var msg2 = []byte("hello world")
	sha2 := sha256.Sum256(msg2)
	print(sha2)
	print("\n")
}