package main

import (
	"bytes"
	"errors"
	"image/png"
	"strings"
	"testing"

	"github.com/OZON08/qrencode/qrencode-go"
)

type failingWriter struct{}

func (failingWriter) Write([]byte) (int, error) {
	return 0, errors.New("write failed")
}

func setLocale(t *testing.T, lcAll, lcMessages, lang string) {
	t.Helper()
	t.Setenv("LC_ALL", lcAll)
	t.Setenv("LC_MESSAGES", lcMessages)
	t.Setenv("LANG", lang)
}

func encodedPNG(t *testing.T, content string) []byte {
	t.Helper()
	grid, err := qrencode.Encode(content, qrencode.ECLevelQ)
	if err != nil {
		t.Fatal(err)
	}
	var buf bytes.Buffer
	if err := png.Encode(&buf, grid.Image(8)); err != nil {
		t.Fatal(err)
	}
	return buf.Bytes()
}

func TestRunJoinsArgsWithSpace(t *testing.T) {
	var stdout bytes.Buffer
	if err := run([]string{"some", "text"}, &stdout); err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(stdout.Bytes(), encodedPNG(t, "some text")) {
		t.Error("stdout does not match the QR code for \"some text\"")
	}
}

func TestRunWithoutArgsWritesPNG(t *testing.T) {
	var stdout bytes.Buffer
	if err := run(nil, &stdout); err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(stdout.Bytes(), encodedPNG(t, "")) {
		t.Error("stdout does not match the QR code for empty content")
	}
}

func TestRunContentTooLarge(t *testing.T) {
	setLocale(t, "C", "", "")
	var stdout bytes.Buffer
	err := run([]string{strings.Repeat("a", 5000)}, &stdout)
	if err == nil || !strings.Contains(err.Error(), "error encoding QR code: Content too large") {
		t.Errorf("err = %v, want the encoding error", err)
	}
	if stdout.Len() != 0 {
		t.Errorf("stdout must stay empty, got %d bytes", stdout.Len())
	}
}

func TestRunWriteError(t *testing.T) {
	setLocale(t, "C", "", "")
	err := run([]string{"text"}, failingWriter{})
	if err == nil || !strings.Contains(err.Error(), "error writing PNG: write failed") {
		t.Errorf("err = %v, want the write error", err)
	}
}

func TestRunLocalizesErrors(t *testing.T) {
	const german, english = "Fehler beim Schreiben der PNG-Datei", "error writing PNG"
	tests := []struct {
		name, lcAll, lcMessages, lang, want string
	}{
		{"LC_ALL German", "de_DE.UTF-8", "", "", german},
		{"LANG German", "", "", "de_DE.UTF-8", german},
		{"LC_ALL overrides LANG", "C", "", "de_DE.UTF-8", english},
		{"LC_MESSAGES overrides LANG", "", "en_US.UTF-8", "de_DE.UTF-8", english},
		{"no locale", "", "", "", english},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setLocale(t, tt.lcAll, tt.lcMessages, tt.lang)
			err := run([]string{"text"}, failingWriter{})
			if err == nil || !strings.HasPrefix(err.Error(), tt.want) {
				t.Errorf("err = %v, want prefix %q", err, tt.want)
			}
		})
	}
}
