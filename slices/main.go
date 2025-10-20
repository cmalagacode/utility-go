package main

import "fmt"

func Example1() {
	// regular array below
	var _ = [5]int{1, 2, 3, 4, 5}
	// slice below
	var _ = []int{1, 2, 3, 4, 5}
	// slice below
	data := make([]int, 5)
	otherSlice := []int{999, 3443, 3443}
	data = append(data, otherSlice...)
	data = append(data, 1000)

	for i := 0; i < len(data); i++ {
		fmt.Println(data[i])
	}
}

func main() {
	Example1()
}
