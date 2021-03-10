package HelloWorld

import (
	"testing"
)

func TestHello(t *testing.T) {
	got := Hello("nathan")
	want := "Hello, nathan"

	if got != want {
		t.Errorf("GOT: %q WANT: %q", got, want)
	}
}

func TestHello2(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("GOT: %q WANT: %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("nathan")
		want := "Hello, nathan"
		assertCorrectMessage(t, got, want)
	})
	t.Run("empty string defaults to 'World'", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
}
