package main

import (
	"errors"
	"fmt"
)

func accessSlice(slice []int, index int) (int, error) {
	if index < 0 || index >= len(slice) {
		return 0, errors.New("Index out of bound")
	}

	return slice[index], nil
}

func main() {
	data := []int{12, 132, 34, 123, 53, 6, 23}

	var indexInput int

	if _, err := fmt.Scan(&indexInput); err != nil {
		fmt.Println("Failed to read the input:", err)
		return
	}

	value, err := accessSlice(data, indexInput)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("item:", indexInput, "value:", value)
}
