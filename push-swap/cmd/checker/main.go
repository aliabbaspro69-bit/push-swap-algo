package main

import (
	"bufio"
	"fmt"
	"os"

	"pushswap/internal/stack"
)

func main() {
	values, err := stack.Parse(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		os.Exit(1)
	}
	if len(values) == 0 {
		return
	}
	m := stack.New(values)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		if err := m.Apply(line); err != nil {
			fmt.Fprintln(os.Stderr, "Error")
			os.Exit(1)
		}
	}
	if m.IsSorted() {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
}
