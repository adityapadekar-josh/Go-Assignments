package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getMostFrequentWords(inputString string) []string {
  countMap := make(map[string]int);
  wordOrder := make([]string, 0);
  maxCount := 0;

  words := strings.Split(strings.TrimSpace(inputString), " ");
  for _, word := range words {
    countMap[word]++;
    wordCount := countMap[word];
      
    if wordCount == 1 {
      wordOrder = append(wordOrder, word);
    }

    maxCount = max(maxCount, wordCount);
  }

  result := make([]string, 0);
  for _, word := range wordOrder {
    if countMap[word] == maxCount {
      result = append(result, word);
    }
  }
  
  return result;
}

func main(){
  scanner := bufio.NewScanner(os.Stdin);
  scanner.Scan();
  inputString := scanner.Text();

  maxCountWord := getMostFrequentWords(inputString);

  fmt.Println(maxCountWord);

}
