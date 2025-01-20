package main

import (
	"fmt"
)

type Quadrilateral interface {
  Area() int
  Perimeter() int
}

type Square struct {
  edge int
}

func NewSquare(edge int) Square {
  return Square{
    edge: edge,
  }
}

func (s Square) Area() int {
  return s.edge * s.edge
}

func (s Square) Perimeter() int {
  return 4 * s.edge
}

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

func Print(quadrilateral Quadrilateral){ 
  fmt.Println("Area =", quadrilateral.Area())
  fmt.Println("Perimeter =", quadrilateral.Perimeter())
}

func main() {
  var choice int
  if _, err := fmt.Scan(&choice); err != nil {
    fmt.Println("Invalid  input")
    return
  }

  rectangle := NewRectangle(20, 10)
  square := NewSquare(15)

  switch choice{
  case 1:
    Print(square)
  case 2:
    Print(rectangle)
  default:
    fmt.Println("Invalid input")
  }

}

