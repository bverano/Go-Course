package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	v := Years(5)
	if v != 35 {
		t.Error("Expected", 35, "Got", v)
	}
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(5)
	}
}

func ExampleYears() {
	fmt.Println(Years(5))
	//Output:
	//35
}

func TestYearsTwo(t *testing.T) {
	v := YearsTwo(10)
	if v != 70 {
		t.Error("Expected", 70, "Got", v)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(10)
	}
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(10))
	//Output:
	//70
}
