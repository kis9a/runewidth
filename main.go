package main

import (
	"fmt"
	"github.com/mattn/go-runewidth"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		fmt.Printf("%d", runewidth.StringWidth(os.Args[1]))
	}
}
