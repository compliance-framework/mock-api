package main

import "testing"

// TestE2ENotifyDeliberateFailure fails on purpose to exercise the Slack
// incident thread (W1-S3-E2E-NOTIFY). Temporary: this PR is never merged.
func TestE2ENotifyDeliberateFailure(t *testing.T) {
	if got, want := 1+1, 3; got != want {
		t.Fatalf("deliberate failure for the notify e2e test: got %d, want %d", got, want)
	}
}
