package main

import (
	"flag"
	"fmt"
	"github.com/gitteri/cpals/pkg/convert"
	"log"
)

var h string
func init() {
	flag.StringVar(&h, "hex", "", "hex to convert to b64")
}

// Challenge 1:
// Convert hex to base64
// Example usage: hextob64 --hex="0123456789abcdef"
func main() {
	flag.Parse()
	if h == "" {
		log.Fatal("must pass hex flag")
	}

	b64, err := convert.HexToB64(h)
	if err != nil {
		log.Fatalf("could not convert hex to base 64: %v", err)
	}

	fmt.Println(b64)
}
