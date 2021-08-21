package greet

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Reid")

	got := buffer.String()
	want := "Hello, Reid"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
