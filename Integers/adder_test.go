package Integers

import (
	"testing"
	"fmt"
)

func TestAdd(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("Expected: '%d' \nGot: '%d'", expected, sum)
	}
}

func ExampleAdd() {
	sum :=  Add(5, 5)
	fmt.Println(sum)
	// Output: 10
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(2,2)
	}
}