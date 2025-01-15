package main

import (
  "fmt"
  "log"
)

var ITEMS = []string{"qwe", "wer", "ert", "rty", "tyu", "yui", "uio", "iop"};
var ITEMS_LEN = len(ITEMS);

func main(){
  var index1, index2 int;

  if _, err := fmt.Scan(&index1, &index2); err != nil {
    log.Fatal("Invalid input");
  }

  if index1 < 0 || index1 >= ITEMS_LEN || index2 < 0 || index2 >= ITEMS_LEN || index1 > index2 {
    log.Fatal("Incorrect Indexes");
  }

  fmt.Println(ITEMS[:index1 + 1]);
  fmt.Println(ITEMS[index1:index2 + 1]);
  fmt.Println(ITEMS[index2:]);
}
