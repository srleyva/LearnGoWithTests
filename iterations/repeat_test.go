package iterations

import (
	"testing"
	"fmt"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 6)
	expected := "aaaaaa"

	if repeated != expected {
		t.Errorf("Expected: %s \nActual: %s", expected, repeated)
	}
}

func ExampleRepeat() {
	fmt.Println(Repeat("a", 4))
	//Output: aaaa
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 4)
	}
}