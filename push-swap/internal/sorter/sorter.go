package sorter

import "pushswap/internal/stack"

type sorter struct {
	m   *stack.Machine
	ops []string
}

func Sort(values []int) []string {
	s := &sorter{m: stack.New(values)}
	for len(s.m.A) > 3 {
		s.do("pb")
	}
	s.sortSmall()
	for len(s.m.B) > 0 {
		s.insertCheapest()
	}
	s.rotateMinToTop()
	return s.ops
}

func (s *sorter) do(op string) {
	_ = s.m.Apply(op)
	s.ops = append(s.ops, op)
}

func (s *sorter) sortSmall() {
	switch len(s.m.A) {
	case 2:
		if s.m.A[0] > s.m.A[1] {
			s.do("sa")
		}
	case 3:
		if s.isMax(0) {
			s.do("ra")
		} else if s.isMax(1) {
			s.do("rra")
		}
		if s.m.A[0] > s.m.A[1] {
			s.do("sa")
		}
	}
}

func (s *sorter) isMax(i int) bool {
	for _, v := range s.m.A {
		if v > s.m.A[i] {
			return false
		}
	}
	return true
}

func (s *sorter) insertCheapest() {
	bestCost, bestA, bestB := -1, 0, 0
	for i, v := range s.m.B {
		ma := signedMoves(s.targetIndex(v), len(s.m.A))
		mb := signedMoves(i, len(s.m.B))
		c := cost(ma, mb)
		if bestCost < 0 || c < bestCost {
			bestCost, bestA, bestB = c, ma, mb
		}
	}
	s.move(bestA, bestB)
	s.do("pa")
}

func (s *sorter) targetIndex(v int) int {
	a := s.m.A
	minIdx, best := 0, -1
	for i, x := range a {
		if x < a[minIdx] {
			minIdx = i
		}
		if x > v && (best < 0 || x < a[best]) {
			best = i
		}
	}
	if best < 0 {
		return minIdx
	}
	return best
}

func (s *sorter) move(ma, mb int) {
	for ma > 0 && mb > 0 {
		s.do("rr")
		ma--
		mb--
	}
	for ma < 0 && mb < 0 {
		s.do("rrr")
		ma++
		mb++
	}
	for ma > 0 {
		s.do("ra")
		ma--
	}
	for ma < 0 {
		s.do("rra")
		ma++
	}
	for mb > 0 {
		s.do("rb")
		mb--
	}
	for mb < 0 {
		s.do("rrb")
		mb++
	}
}

func (s *sorter) rotateMinToTop() {
	a := s.m.A
	minIdx := 0
	for i, v := range a {
		if v < a[minIdx] {
			minIdx = i
		}
	}
	k := signedMoves(minIdx, len(a))
	for ; k > 0; k-- {
		s.do("ra")
	}
	for ; k < 0; k++ {
		s.do("rra")
	}
}

func signedMoves(index, length int) int {
	if index <= length/2 {
		return index
	}
	return index - length
}

func cost(ma, mb int) int {
	if ma*mb > 0 {
		if abs(ma) > abs(mb) {
			return abs(ma)
		}
		return abs(mb)
	}
	return abs(ma) + abs(mb)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
