package main

import (
  "errors"
  "fmt"
  "log"
)


var romanSymbols = map[string]int{
  "I" : 1,
  "V" : 5,
  "X" : 10,
  "L" : 50,
  "C" : 100,
  "D" : 500,
  "M" : 1000,
}

func convertRomanNumeralToIntegerRepresentation(romanNumeral string) (int, error) {
  integerRepresentation := 0;

  for i := 0; i < len(romanNumeral); i++ {
    currentSymbol := string(romanNumeral[i]);
    currentValue, exists := romanSymbols[currentSymbol]
    if !exists {
      return 0, errors.New("Invalid Roman numeral representation");
    }
    
    if i + 1 < len(romanNumeral) {
      nextSymbol := string(romanNumeral[i + 1]);
      nextValue, exists := romanSymbols[nextSymbol];
      if !exists {
        return 0, errors.New("Invalid Roman numeral representation");
      }

      if currentValue < nextValue {
        integerRepresentation += nextValue - currentValue;
        i++;
        continue;
      }
    }

    integerRepresentation += currentValue;
  }

  return integerRepresentation, nil;
}


func main(){
  var romanNumber string;

  if _, err := fmt.Scan(&romanNumber); err != nil {
    log.Fatal("Failed to read the input");
  }

  integerRepresentation, err := convertRomanNumeralToIntegerRepresentation(romanNumber);
  if err != nil {
    log.Fatal(err.Error());
  }

  fmt.Println("Integer representation of ", romanNumber, " is ", integerRepresentation);
}
