package main

import (
	"bufio"
	"fmt"
	"os"

	"pushswap/internal/sorter"
	"pushswap/internal/stack"
)

func main() {
	values, err := stack.Parse(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		os.Exit(1)
	}
	out := bufio.NewWriter(os.Stdout)
	for _, op := range sorter.Sort(values) {
		fmt.Fprintln(out, op)
	}
	out.Flush()
}
