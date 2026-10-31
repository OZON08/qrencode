package main

import (
	"fmt"
	"image/png"
	"io"
	"os"
	"strings"

	"github.com/OZON08/qrencode/qrencode-go"
)

func main() {
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr))
}

// run encodes args, joined by spaces, as a PNG QR code and returns the exit code.
func run(args []string, stdout, stderr io.Writer) int {
	grid, err := qrencode.Encode(strings.Join(args, " "), qrencode.ECLevelQ)
	if err != nil {
		fmt.Fprintln(stderr, "error encoding QR code:", err)
		return 1
	}
	if err := png.Encode(stdout, grid.Image(8)); err != nil {
		fmt.Fprintln(stderr, "error writing PNG:", err)
		return 1
	}
	return 0
}
