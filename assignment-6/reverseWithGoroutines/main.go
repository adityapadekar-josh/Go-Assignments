package main

import (
	"fmt"
	"runtime"
	"sync"
)

func reverseString(str string, wg *sync.WaitGroup) {
	defer wg.Done()
	var reversedStr = ""

	for i := len(str) - 1; i >= 0; i-- {
		reversedStr += string(str[i])
	}

	fmt.Println(reversedStr)
	fmt.Println(runtime.NumGoroutine())

}

func main() {
	var inputString string
	var wg sync.WaitGroup

	if _, err := fmt.Scan(&inputString); err != nil {
		fmt.Println("Failed to read input")
		return
	}

	wg.Add(1)

	go reverseString(inputString, &wg)

	wg.Wait()
}
