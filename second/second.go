package main

import "fmt"

func secondToLast(values []int) int {
	return values[len(values)-2]
}

func safeSecondToLast(values []int) (n int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %v", r)
		}
	}()

	return secondToLast(values), nil
}

func main() {
	nums := []int{10, 20, 30, 40, 50}
	fmt.Println(secondToLast(nums)) // 40
	fmt.Println(safeSecondToLast(nil))
	print("Done")
}
