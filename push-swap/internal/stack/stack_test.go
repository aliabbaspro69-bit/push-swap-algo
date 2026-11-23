package stack

import "testing"

func TestParse(t *testing.T) {
	values, err := Parse([]string{"3 1", "2"})
	if err != nil || len(values) != 3 || values[0] != 3 || values[2] != 2 {
		t.Fatalf("unexpected result: %v %v", values, err)
	}
	if _, err := Parse([]string{"1 1"}); err == nil {
		t.Fatal("duplicate accepted")
	}
	if _, err := Parse([]string{"1 two"}); err == nil {
		t.Fatal("non integer accepted")
	}
}

func TestApply(t *testing.T) {
	m := New([]int{2, 1, 3})
	for _, op := range []string{"sa", "pb", "ra", "pa"} {
		if err := m.Apply(op); err != nil {
			t.Fatal(err)
		}
	}
	if len(m.B) != 0 || m.A[0] != 1 || m.A[1] != 3 || m.A[2] != 2 {
		t.Fatalf("A=%v B=%v", m.A, m.B)
	}
	if m.Apply("xx") == nil {
		t.Fatal("unknown instruction accepted")
	}
}
