package saying

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	s := Greet("James")

	if s != "Hello, Welcome James" {
		t.Error("got", s, "expected", "Hello, Welcome James")
	}
}

func ExampleGreet() {
	fmt.Println(Greet("James"))
	// output:
	// Hello, Welcome James
}

func BenchmarkGreet(b *testing.B) {
	for i :=0; i < b.N; i++ {
		Greet("James")
	}
}