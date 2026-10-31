package main

import (
	"fmt"
	"image/png"
	"io"
	"os"
	"strings"

	"github.com/OZON08/qrencode/qrencode-go"
)

// messages holds the user-facing error prefixes of one language.
type messages struct {
	encode, write string
}

var (
	english = messages{"error encoding QR code", "error writing PNG"}
	german  = messages{"Fehler beim Erzeugen des QR-Codes", "Fehler beim Schreiben der PNG-Datei"}
)

// localMessages selects the language from the first locale variable that is
// set, following the POSIX precedence, and falls back to English.
func localMessages() messages {
	for _, name := range []string{"LC_ALL", "LC_MESSAGES", "LANG"} {
		if value := os.Getenv(name); value != "" {
			if strings.HasPrefix(value, "de") {
				return german
			}
			return english
		}
	}
	return english
}

func main() {
	if err := run(os.Args[1:], os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// run encodes args, joined by spaces, as a QR code and writes it to w as PNG.
func run(args []string, w io.Writer) error {
	text := localMessages()
	grid, err := qrencode.Encode(strings.Join(args, " "), qrencode.ECLevelQ)
	if err != nil {
		return fmt.Errorf("%s: %w", text.encode, err)
	}
	if err := png.Encode(w, grid.Image(8)); err != nil {
		return fmt.Errorf("%s: %w", text.write, err)
	}
	return nil
}
