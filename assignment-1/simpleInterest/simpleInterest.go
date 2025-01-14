/*
Q] Write a program to calculate the simple interest
First-line has the comma-separated values of the Principal, rate and time (in years) respective

*constraints: Round simple interest float value to 2 decimal places
*/

package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func calculateSimpleInterest(principal, rate, time float64) float64 {
  return (principal * rate * time) / 100;
}


func main(){
  var inputString string;

  if _, err := fmt.Scan(&inputString); err != nil {
    log.Fatal("Failed to accept input");
  }

  inputValues := strings.Split(inputString, ",");

  principalAmount, err1 := strconv.ParseFloat(inputValues[0], 64);
  interestRate, err2 := strconv.ParseFloat(inputValues[1], 64);
  timePeriod, err3 := strconv.ParseFloat(inputValues[2], 64);

  if err1 != nil || err2 != nil || err3 != nil {
    log.Fatal("Invalid input string")
  }

  simpleInterest := calculateSimpleInterest(principalAmount, interestRate, timePeriod);
  
  fmt.Printf("Simple interest: %.2f\n", simpleInterest);
}
