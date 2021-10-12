package bubble

import "testing"

func Test_BubbleSort(t *testing.T) {
	values := []int{4, 1, 7, 9, 2, 5, 3}
	BubbleSort(values)
	if values[0] != 1 || values[1] != 2 || values[2] != 3 || values[3] != 4 ||
		values[4] != 5 || values[5] != 7 || values[6] != 9 {
		t.Error("BubbleSort() failed. Got", values, "Expected 1 2 3 4 5 7 9")
	}
}

func Test_BubbleSort2(t *testing.T) {
	values := []int{5, 4, 3, 2, 1}
	BubbleSort(values)
	if values[0] != 1 || values[1] != 2 || values[2] != 3 || values[3] != 4 ||
		values[4] != 5 {
		t.Error("QuickSort() failed. Got", values, "Expected 1 2 3 4 5")
	}
}

func Test_BubbleSort3(t *testing.T) {
	values := []int{1, 2, 3, 4, 5}
	BubbleSort(values)
	if values[0] != 1 || values[1] != 2 || values[2] != 3 || values[3] != 4 ||
		values[4] != 5 {
		t.Error("QuickSort() failed. Got", values, "Expected 1 2 3 4 5")
	}
}
