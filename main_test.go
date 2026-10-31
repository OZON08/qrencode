package main

import (
	"bytes"
	"errors"
	"image/png"
	"strings"
	"testing"
)

type failingWriter struct{}

func (failingWriter) Write([]byte) (int, error) {
	return 0, errors.New("write failed")
}

func assertValidPNG(t *testing.T, args []string) {
	t.Helper()
	var stdout, stderr bytes.Buffer
	if code := run(args, &stdout, &stderr); code != 0 {
		t.Fatalf("exit code = %d, stderr = %q", code, stderr.String())
	}
	if stderr.Len() != 0 {
		t.Errorf("unexpected stderr: %q", stderr.String())
	}
	if _, err := png.Decode(&stdout); err != nil {
		t.Errorf("stdout is not a valid PNG: %v", err)
	}
}

func TestRunWritesPNG(t *testing.T) {
	assertValidPNG(t, []string{"some", "text"})
}

func TestRunWithoutArgsWritesPNG(t *testing.T) {
	assertValidPNG(t, nil)
}

func TestRunContentTooLarge(t *testing.T) {
	var stdout, stderr bytes.Buffer
	if code := run([]string{strings.Repeat("a", 5000)}, &stdout, &stderr); code != 1 {
		t.Errorf("exit code = %d, want 1", code)
	}
	if stdout.Len() != 0 {
		t.Errorf("stdout must stay empty, got %d bytes", stdout.Len())
	}
	if !strings.Contains(stderr.String(), "Content too large") {
		t.Errorf("stderr = %q, want the encoding error", stderr.String())
	}
}

func TestRunWriteError(t *testing.T) {
	var stderr bytes.Buffer
	if code := run([]string{"text"}, failingWriter{}, &stderr); code != 1 {
		t.Errorf("exit code = %d, want 1", code)
	}
	if !strings.Contains(stderr.String(), "error writing PNG") {
		t.Errorf("stderr = %q, want the write error", stderr.String())
	}
}
