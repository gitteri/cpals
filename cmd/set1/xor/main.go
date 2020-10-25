package main

import (
	"fmt"
	"github.com/gitteri/cpals/pkg/bits"
	"log"
	"os"
)

// Challenge 2:
// xor two hex strings
// Example usage: xor 1ef 2ef
func main() {
	hexStrs := os.Args[1:]
	if len(hexStrs) != 2 {
		log.Fatalf("can only xor 2 strings. provided %d", len(hexStrs))
	}

	xor, err := bits.XorHexStrs(hexStrs[0], hexStrs[1])
	if err != nil {
		log.Fatalf("cannot xor byte arrays: %v", err)
	}

	fmt.Println(xor)
}
