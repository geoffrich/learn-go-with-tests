// one package per directory
package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

// Example -- executed just like tests
// Better than examples living in documentation because it won't get out of date
// Automatically added to godoc documentation
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// This function is only run during tests because of the below comment
	// Output: 6
}
