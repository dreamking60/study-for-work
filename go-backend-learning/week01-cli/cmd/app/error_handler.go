package main

import (
	"fmt"
	"os"
)

func handleError(err error) {
	fmt.Fprintf(os.Stderr, "level=ERROR msg=%q error=%q\n", "application failed", err.Error())
	os.Exit(1)
}
