package main

// Testing package built into language
import "testing"

/*
Writing a test:
1. File is named xxx_test.go
2. Test function starts with word "Test"
3. Test function takes one argument: t *testing.T
*/

func TestHello(t *testing.T) { // t *testing.T -- our "hook" into the testing framework

	// functions can be assigned to variables
	// when you have multiple consecutive parameters of the same type, you can omit type names
	assertCorrectMessage := func(t *testing.T, got, want string) {
		// tells the test suite that this method is a helper
		// failures are reported in function call, not test helper
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	// subtest
	t.Run("saying hello to people", func(t *testing.T) {
		// How can it get hello -- is it because they're in the same package?
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Elodie", "French")
		want := "Bonjour, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in German", func(t *testing.T) {
		got := Hello("Elodie", "German")
		want := "Hallo, Elodie"
		assertCorrectMessage(t, got, want)
	})

}
