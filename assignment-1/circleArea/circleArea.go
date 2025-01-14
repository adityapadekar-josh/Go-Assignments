/* 
Q] Write a program to calculate the area of the circle, First line has a value of the radius of the circle

*constraint
1. Use const PI from the package math package
2. Use the Pow function from the math package
3. Round area float value to 2 decimal places
*/ 

package main

import (
	"fmt"
	"log"
	"math"
)

func calculateAreaOfCircle(radius float64) float64 {
  return math.Pi * math.Pow(radius, 2);
}

func main(){
  var radius float64;

  if _, err := fmt.Scan(&radius); err != nil {
    log.Fatal("Failed to accept input");
  }

  circleArea := calculateAreaOfCircle(radius);

  fmt.Printf("Area of circle is %.2f\n", circleArea);
}
