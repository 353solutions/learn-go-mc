package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

func uptime(fname string) (time.Duration, error) {
	data, err := os.ReadFile(fname)
	if err != nil {
		return 0, err
	}

	s := strings.TrimSpace(string(data))
	t, err := time.Parse("2006-01-02T15:04:05", s)
	if err != nil {
		return 0, fmt.Errorf("%q: parse - %w", fname, err)
	}

	return time.Since(t), nil
}

func main() {
	d, err := uptime("start.txt")
	if err != nil {
		// errors.Is checks if any error in the chain matches a target value.
		// Useful for sentinel errors like os.ErrNotExist.
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("start file not found")
			return
		}

		// errors.AsType is a generic helper that combines the type assertion and As call.
		// Useful for structured errors like *os.PathError which carries Path and Op fields.
		if pathErr, ok := errors.AsType[*os.PathError](err); ok {
			fmt.Printf("file error: op=%s path=%s\n", pathErr.Op, pathErr.Path)
			return
		}

		fmt.Println("error:", err)
		return
	}

	fmt.Println("Server has been running for:", d)
}
