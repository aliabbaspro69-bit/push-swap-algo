package stack

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalid = errors.New("invalid")

func Parse(args []string) ([]int, error) {
	seen := make(map[int]bool)
	values := []int{}
	for _, arg := range args {
		for _, field := range strings.Fields(arg) {
			n, err := strconv.Atoi(field)
			if err != nil || seen[n] {
				return nil, ErrInvalid
			}
			seen[n] = true
			values = append(values, n)
		}
	}
	return values, nil
}
