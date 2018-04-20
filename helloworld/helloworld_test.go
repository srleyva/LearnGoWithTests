package main

import (
	"testing"
	"fmt"
)

func TestHello(t *testing.T) {

	t.Run("say hello to Stephen", func(t *testing.T) {
		got := Hello("Stephen")
		want := "Hello, Stephen!"

		if got != want {
			t.Errorf("Got: %s, \nWant: %s", got, want)
		}
	})

	t.Run("say hello world when empty string applied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World!"
		if got != want {
			t.Errorf("Got: %s, \nWant: %s", got, want)
		}
	})

}

func ExampleHello() {
	fmt.Println(Hello("Stephen"))
	//Output: Hello, Stephen!
}

func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Hello("Stephen")
	}
}