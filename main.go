package main

import (
	"fmt"
	"os"
)

var Version = "(development)"

func main() {
	fmt.Fprintf(os.Stderr, "%s - Version %s\n", os.Args[0], Version)
}
