package grokalgorithms

import "testing"

func TestSimpleSearchFirstItem(t *testing.T) {
	want := 0
	got := SimpleSearch([]string{"One", "Two", "Three"}, "One")
	if got != want {
		t.Errorf("Simple search got: %d, want: %d.", got, want)
	}
}

func TestSimpleSearchMiddleItem(t *testing.T) {
	want := 1
	got := SimpleSearch([]string{"One", "Two", "Three"}, "Two")
	if got != want {
		t.Errorf("Simple search got: %d, want: %d.", got, want)
	}
}

func TestSimpleSearchNoMatch(t *testing.T) {
	want := -1
	got := SimpleSearch([]string{"One", "Two", "Three"}, "Sixty")
	if got != want {
		t.Errorf("A Simple Search with no match should return %d, got: %d.", want, got)
	}
}
