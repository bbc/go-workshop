package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
)

func main() {
	i := 0
	var file string
	for i, file = range os.Args[1:] {
		compress(file)
	}
	fmt.Printf("Compressed %d files\n", i+1)
}

func compress(filename string) error {
	in, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(filename + ".gz")
	if err != nil {
		return err
	}
	defer out.Close()

	gzout := gzip.NewWriter(out)
	_, err = io.Copy(gzout, in)
	gzout.Close()

	return err
}
