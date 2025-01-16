package main

import (
	"fmt"
	"log"
)

type hello struct {
	message string
}

func acceptAnything(value interface{}) {
	switch valueType := value.(type) {
	case int:
		fmt.Println("This is a value of type Integer,", valueType)
	case string:
		fmt.Println("This is a value of type String,", valueType)
	case bool:
		fmt.Println("This is a value of type Boolean,", valueType)
	case hello:
		fmt.Println("This is a value of type Hello,", valueType.message)
	default:
		fmt.Println("Unsupported data type")
	}
}

func main() {
	var choice int
  if _, err := fmt.Scanln(&choice); err != nil {
    log.Fatal("Failed to read input")
  }

	switch choice {
	case 1:
		acceptAnything(42)
	case 2:
		acceptAnything("Hello World!") 
	case 3:
		acceptAnything(true) 
	case 4:
		acceptAnything(hello{"Greetings from Hello class"}) 
	default:
		fmt.Println("Invalid choice.")
	}
}

