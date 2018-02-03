package grokalgorithms

import "testing"

func TestBinaryChopMiddleItem(t *testing.T) {
	want := 1
	got := BinaryChop([]int{1, 2, 3}, 2)
	if got != want {
		t.Errorf("Binary Chop got: %d, want: %d.", got, want)
	}
}

func TestBinaryChopLeftItem(t *testing.T) {
	want := 0
	got := BinaryChop([]int{1, 2, 3}, 1)
	if got != want {
		t.Errorf("Binary Chop got: %d, want: %d.", got, want)
	}
}

func TestBinaryChopRightItem(t *testing.T) {
	want := 2
	got := BinaryChop([]int{1, 2, 3}, 3)
	if got != want {
		t.Errorf("Binary Chop got: %d, want: %d.", got, want)
	}
}

func TestBinaryChopFiveItems(t *testing.T) {
	want := 3
	got := BinaryChop([]int{0, 1, 2, 3, 4}, 3)
	if got != want {
		t.Errorf("Binary Chop got: %d, want: %d.", got, want)
	}
}

func TestBinaryChopNotFound(t *testing.T) {
	want := -1
	got := BinaryChop([]int{0, 1, 2}, 42)
	if got != want {
		t.Errorf("Binary Chop not found should return %d, got: %d", want, got)
	}
}

func TestBinaryChopTenItems(t *testing.T) {
	want := 0
	got := BinaryChop([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 0)
	if got != want {
		t.Errorf("Binary Chop got: %d, want: %d.", got, want)
	}
}

func TestBinaryChopElevenItems(t *testing.T) {
	want := 10
	got := BinaryChop([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10)
	if got != want {
		t.Errorf("Binary Chop got: %d, want: %d.", got, want)
	}
}
