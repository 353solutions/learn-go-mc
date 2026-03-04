package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"strings"
)

func fileSize(filename string) (int64, error) {
	// cat road.txt.gz
	f, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	var r io.Reader = f
	if strings.HasSuffix(filename, ".gz") {
		// | gunzip
		gz, err := gzip.NewReader(f)
		if err != nil {
			return 0, err
		}
		defer gz.Close()
		r = gz
	}

	// | wc -c
	return io.Copy(io.Discard, r)
}

func main() {
	size, err := fileSize("road.txt.gz")
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("uncompressed size: %d bytes\n", size)
	fmt.Println(fileSize("road.txt"))
}
