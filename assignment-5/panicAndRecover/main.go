package main

import (
	"fmt"
	"log"
)

func accessSlice(slice []int, index int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("internal error:", r)
			// debug.PrintStack()
		}
	}()

	println("item:", index, "value:", slice[index])
}

func main() {
	data := []int{12, 132, 34, 123, 53, 6, 23}

	var indexInput int

	if _, err := fmt.Scan(&indexInput); err != nil {
		log.Fatal("Failed to read the input:", err)
	}

	accessSlice(data, indexInput)
}
