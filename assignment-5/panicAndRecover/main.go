package main

import (
	"fmt"
)

func accessSlice(slice []int, index int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("internal error:", r)
		}
	}()

	fmt.Println("item:", index, "value:", slice[index])
}

func main() {
	data := []int{12, 132, 34, 123, 53, 6, 23}

	var indexInput int

	if _, err := fmt.Scan(&indexInput); err != nil {
		fmt.Println("Failed to read the input:", err)
		return
	}

	accessSlice(data, indexInput)
}
