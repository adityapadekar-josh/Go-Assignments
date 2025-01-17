package main

import (
  "fmt"
  "log"
)

var Items = []string{"qwe", "wer", "ert", "rty", "tyu", "yui", "uio", "iop"};
var ItemsLen = len(Items);

func main(){
  var index1, index2 int;

  if _, err := fmt.Scan(&index1, &index2); err != nil {
    log.Fatal("Invalid input");
  }

  if index1 < 0 || index1 >= ItemsLen || index2 < 0 || index2 >= ItemsLen || index1 > index2 {
    log.Fatal("Incorrect Indexes");
  }

  fmt.Println(Items[:index1 + 1]);
  fmt.Println(Items[index1:index2 + 1]);
  fmt.Println(Items[index2:]);
}
