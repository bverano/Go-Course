package say

import (
	"fmt"
	"testing"
)

func TestSayHi(t *testing.T) {
	s := SayHi("Bruno")
	if s != "Hi, Bruno. Welcome!" {
		t.Error("Expected", "Hi, Bruno. Welcome", "Got", s)
	}
}

func ExampleSayHi() {
	fmt.Println(SayHi("Bruno"))
	//Output:
	//Hi, Bruno. Welcome!
}

func BenchmarkSayHi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SayHi("Bruno")
	}
}
