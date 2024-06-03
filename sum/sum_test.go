package sum

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("with size 5", func(t *testing.T) {
		array := []int{1, 2, 3, 4, 5}
		got := Sum(array)
		want := 15
		if got != want {
			t.Errorf("Got %v, want %v when given %v", got, want, array)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2, 3}, []int{0, 0, 0, 10})
	want := []int{6, 10}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v and wanted %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(got []int, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v and wanted %v", got, want)
		}
	}
	t.Run("With a tail", func(t *testing.T) {
		got := SumAllTails([]int{999, 2, 2}, []int{-2, 4, 5})
		want := []int{4, 9}
		checkSums(got, want)
	})
	t.Run("Without a tail", func(t *testing.T) {
		got := SumAllTails([]int{15}, []int{-24})
		want := []int{0, 0}
		checkSums(got, want)
	})
	t.Run("Safe sum (empty slice)", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{})
		want := []int{0, 0}
		checkSums(got, want)
	})
}
