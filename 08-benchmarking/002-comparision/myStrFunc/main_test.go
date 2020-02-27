package myStrFunc

import (
	"fmt"
	"strings"
	"testing"
)

const s = `"It is a long established fact that a reader will be distracted by the 
readable content of a page when looking at its layout. The point of using Lorem 
Ipsum is that it has a more-or-less normal distribution of letters, as opposed 
to using 'Content here, content here', making it look like readable English. Many
 desktop publishing packages and web page editors now use Lorem Ipsum as their 
 default model text, and a search for 'lorem ipsum' will uncover many web sites
  still in their infancy. Various versions have evolved over the years, sometimes by
   accident, sometimes on purpose (injected humour and the like)"`

var xs []string

func TestConCat(t *testing.T) {
	s := "I forget that i am ConCat Func"
	xs := strings.Split(s, " ")
	s = ConCat(xs)

	if s != "I forget that i am ConCat Func" {
		t.Error("got", s, "wanted", "I forget that i am ConCat Func")
	}
}

func TestJoin(t *testing.T) {
	s := "I forget that i am Join Func"
	xs := strings.Split(s, " ")
	s = Join(xs)
	if s != "I forget that i am Join Func" {
		t.Error("got", s, "wanted", "I forget that i am Join Func")
	}
}

func BenchmarkConCat(b *testing.B) {
	xs = strings.Split(s, " ")
	for i := 0; i < b.N; i++ {
		ConCat(xs)
	}
}

func BenchmarkJoin(b *testing.B) {
	xs = strings.Split(s, " ")
	for i := 0; i < b.N; i++ {
		Join(xs)
	}
}

func ExampleConCat() {
	s := "I forget that i am ConCat Func"
	xs := strings.Split(s, " ")
	fmt.Println(ConCat(xs))
	// Output:
	// I forget that i am ConCat Func
}

func ExampleJoin() {
	s := "I forget that i am Join Func"
	xs := strings.Split(s, " ")
	fmt.Println(Join(xs))
	// Output:
	// I forget that i am Join Func
}
