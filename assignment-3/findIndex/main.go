package main

import (
  "fmt"
  "log"
)

var indexToDayMapping = map[int]string{
  1 : "Monday",
  2 : "Tuesday",
  3 : "Wednesday",
  4 : "Thursday",
  5 : "Friday",
  6 : "Saturday",
  7 : "Sunday",
};

func getDayByIndex(index int) string {
  day, ok := indexToDayMapping[index];

  if ok {
    return day;
  }
  return "Not A Day";
}

func main()  {
  var indexInput int; 

  if _, err := fmt.Scan(&indexInput); err != nil {
    log.Fatal("Invalid input index");
  }

  dayResult := getDayByIndex(indexInput);

  fmt.Println("Corresponding day for index", indexInput, "is", dayResult);
}
