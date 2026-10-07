package version

import "testing"

func TestHello(t *testing.T) {
	const want = "hello from mock-api 0.1.0"
	if got := Hello(); got != want {
		t.Fatalf("Hello() = %q, want %q (update this test when bumping Version)", got, want)
	}
}
