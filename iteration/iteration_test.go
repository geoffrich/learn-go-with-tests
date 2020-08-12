package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 6)
	expected := "aaaaaa"

	// do we always have to write our own assertions?
	// or is there a built in Assert?
	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

// benchmark
// runs a function N times -- Go determines what this value should be for best results
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeated := Repeat("b", 5)
	fmt.Println(repeated)
	// Output: bbbbb
}
