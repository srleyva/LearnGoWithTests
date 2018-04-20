package Arrays

import (
	"testing"
	"fmt"
	"reflect"
)

func CheckSum(t *testing.T, actual, expected []int) {
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got %v want %v", actual, expected)
	}
}
func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1,2,3}

		actual := Sum(numbers)
		expected := 6

		if actual != expected {
			t.Errorf("Got: %d \nExpected: %d", actual, expected)
		}
	})

}

func TestSumAll(t *testing.T) {
	t.Run("sum multiple lists", func(t *testing.T) {
		actual := SumAll([]int{1,2}, []int{2,2,3})
		expected := []int{3, 7}
		CheckSum(t, actual, expected)
	})
}

func TestSumAllTails(t *testing.T) {

	t.Run("Test summing of all tails", func(t *testing.T) {
		actual := SumAllTails([]int{2,3,4}, []int{2,5,6})
		expected := []int{7, 11}
		CheckSum(t, actual, expected)
	})

	t.Run("Safely pass in empty slice", func(t *testing.T) {
		actual := SumAllTails([]int{}, []int{1,2})
		expected := []int{0, 2}
		CheckSum(t, actual, expected)
	})
}

func ExampleSum() {
	numbers := []int{1,2,3,4,5}
	fmt.Println(Sum(numbers))
	//Output: 15
}

func ExampleSumAll() {
	fmt.Println(SumAll([]int{1,2}, []int{2,2,3}))
	// Output: [3 7]
}

func ExampleSumAllTails() {
	fmt.Println(SumAllTails([]int{2, 4, 6}, []int{0, 1, 0}))
	// Output: [10 1]
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		numbers := []int{1,2,3,4,5}
		Sum(numbers)
	}
}

func BenchmarkSumAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumAll([]int{1,2}, []int{2,2,3})
	}
}

func BenchmarkSumAllTails(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumAllTails([]int{2, 4, 6}, []int{0, 1, 0})
	}
}