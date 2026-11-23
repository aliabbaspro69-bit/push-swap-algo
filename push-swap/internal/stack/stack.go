package stack

type Machine struct {
	A []int
	B []int
}

func New(values []int) *Machine {
	a := make([]int, len(values))
	copy(a, values)
	return &Machine{A: a, B: []int{}}
}

func (m *Machine) Apply(op string) error {
	switch op {
	case "sa":
		swap(m.A)
	case "sb":
		swap(m.B)
	case "ss":
		swap(m.A)
		swap(m.B)
	case "pa":
		m.A, m.B = push(m.A, m.B)
	case "pb":
		m.B, m.A = push(m.B, m.A)
	case "ra":
		rotate(m.A)
	case "rb":
		rotate(m.B)
	case "rr":
		rotate(m.A)
		rotate(m.B)
	case "rra":
		reverseRotate(m.A)
	case "rrb":
		reverseRotate(m.B)
	case "rrr":
		reverseRotate(m.A)
		reverseRotate(m.B)
	default:
		return ErrInvalid
	}
	return nil
}

func (m *Machine) IsSorted() bool {
	if len(m.B) > 0 {
		return false
	}
	for i := 1; i < len(m.A); i++ {
		if m.A[i-1] > m.A[i] {
			return false
		}
	}
	return true
}

func swap(s []int) {
	if len(s) >= 2 {
		s[0], s[1] = s[1], s[0]
	}
}

func rotate(s []int) {
	if len(s) >= 2 {
		first := s[0]
		copy(s, s[1:])
		s[len(s)-1] = first
	}
}

func reverseRotate(s []int) {
	if len(s) >= 2 {
		last := s[len(s)-1]
		copy(s[1:], s[:len(s)-1])
		s[0] = last
	}
}

func push(dst, src []int) ([]int, []int) {
	if len(src) == 0 {
		return dst, src
	}
	return append([]int{src[0]}, dst...), src[1:]
}
