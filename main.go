package main

import (
    "fmt"
    "os"
	"github.com/matthieuEv/go-sha256/sha256"
)

func main() {
	if len(os.Args) > 1{
		for i:=1; i<len(os.Args); i++ {
			fmt.Println(os.Args[i],": ",sha256.Sum256([]byte(os.Args[i])))
		}
	} else {
		fmt.Println("  No arguments passed")
	}
}