package main

import (
	"bytes"
	"fmt"
	"image/png"
	"os"

	"github.com/d3vilh/qrencode/qrencode-go"
)

func main() {
	var buf bytes.Buffer
	for i, arg := range os.Args {
		if i > 1 {
			if err := buf.WriteByte(' '); err != nil {
				panic(err)
			}
		}
		if i > 0 {
			if _, err := buf.WriteString(arg); err != nil {
				panic(err)
			}
		}
	}
	grid, err := qrencode.Encode(buf.String(), qrencode.ECLevelQ)
	if err != nil {
		panic(err)
	}
	if err := png.Encode(os.Stdout, grid.Image(8)); err != nil {
		fmt.Fprintln(os.Stderr, "error writing PNG:", err)
		os.Exit(1)
	}
}
