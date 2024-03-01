package main

import (
	"fmt"

	// pull this in with `go mod tidy`
	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())
}
