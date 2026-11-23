package sorter

import (
	"math/rand"
	"testing"

	"pushswap/internal/stack"
)

func TestSortSorts(t *testing.T) {
	for _, size := range []int{0, 1, 2, 3, 4, 5, 10, 100} {
		for trial := 0; trial < 20; trial++ {
			values := rand.Perm(size)
			m := stack.New(values)
			for _, op := range Sort(values) {
				if err := m.Apply(op); err != nil {
					t.Fatal(err)
				}
			}
			if !m.IsSorted() {
				t.Fatalf("size %d not sorted: %v", size, values)
			}
		}
	}
}
