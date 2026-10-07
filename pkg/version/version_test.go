package version

import "testing"

func TestHello(t *testing.T) {
	want := "hello from mock-api " + Version
	if got := Hello(); got != want {
		t.Fatalf("Hello() = %q, want %q", got, want)
	}
}
