package main

import (
	"io"
	"os"
)

func echo(in io.Reader, out io.Writer) {
	io.Copy(out, in)
}

func main() {
	// ADD CODE call to echo here...
	// ADD sleep ...
	os.Exit(0)
}
