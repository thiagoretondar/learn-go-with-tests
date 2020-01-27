package helloworld

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Thiago")
	want := "Hello, Thiago"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
