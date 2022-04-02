package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	src, err := readInput()

	if err != nil {
		fail(err)
	}

	wordCount := 0

	trimmed := strings.TrimSpace(src)

	if (trimmed == "") {
		fmt.Println(wordCount)
		return
	}

	for _, symbol := range trimmed {
		if (symbol == ' ') {
			wordCount++;
		}
	}

	fmt.Println(wordCount + 1)
}

// readInput reads pattern and source string
// from command line arguments and returns them.
func readInput() (src string, err error) {
	flag.Parse()
	src = strings.Join(flag.Args(), "")
	return src, nil
}

// fail prints the error and exits.
func fail(err error) {
	fmt.Println("Error:", err)
	os.Exit(1)
}