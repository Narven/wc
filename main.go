package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	// defining a boolean flag -l to count lines instead of words
	lines := flag.Bool("l", false, "Count lines")
	// parsing the flags provided by the user
	flag.Parse()
	/// calling the count function to count the number of words (or lines)
	// received from the standard input and printing it out
	fmt.Println(count(os.Stdin, *lines))
}

func count(r io.Reader, countLines bool) int {
	// a scanner is used to read text from a Reader (such as files)
	scanner := bufio.NewScanner(r)

	// if the count lines flag is not set, we want to count words so we defined
	// the scanned split type to words (default is split by lines)
	if !countLines {
		// define the scanner split type to words (default is split by lines)
		scanner.Split(bufio.ScanWords)
	}

	// defining a counter
	wc := 0

	// for every word scanned, increment the counter
	for scanner.Scan() {
		wc++
	}

	// return the total
	return wc
}
