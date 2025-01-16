package main

import (
	"fmt"
	"log"
)

type Rectangle struct {
    length int
    width  int
}

func NewRectangle(length, width int) Rectangle {
    return Rectangle{
        length: length,
        width:  width,
    }
}

func (r Rectangle) Area() int {
  return r.length * r.width
}

func (r Rectangle) Perimeter() int {
  return 2 * (r.length + r.width)
}

func main() {
  var inputLength, inputWidth int

  if _, err := fmt.Scan(&inputLength, &inputWidth); err != nil {
    log.Fatal("Invalid  input")
  }

  rectangle := NewRectangle(inputLength, inputWidth)

  fmt.Println("Area =", rectangle.Area())
  fmt.Println("Perimenter =", rectangle.Perimeter())
}

