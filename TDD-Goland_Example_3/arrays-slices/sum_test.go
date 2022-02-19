package slices

import (
	"fmt"
	"reflect"
	"testing"
)

var numbers = []int{2, 3, 4, 6, 9, 11}

func TestSum(t *testing.T) {
	t.Run("a collection of 5 numbers", func(t *testing.T) {

		got := Sum(numbers)
		expected := 35

		if got != expected {
			t.Errorf("Expected: '%d', Got: '%d'", expected, got)
		}
	})

	t.Run("a collection of any amount", func(t *testing.T) {
		var numbers = []int{1, 2, 3}
		got := Sum(numbers)
		expected := 6

		if got != expected {
			t.Errorf("Expected: '%d', Got: '%d'", expected, got)
		}
	})

}

func ExampleSum() {
	result := Sum(numbers)
	fmt.Println(result)

	//Output: 35
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum(numbers)
	}
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Want: '%v', Got: '%v'", want, got)
	}
}

func TestSumAllV2(t *testing.T) {
	got := SumAllV2([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Want: '%v', Got: '%v'", want, got)
	}
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sums of tails of", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSums(t, got, want)
	})

}
