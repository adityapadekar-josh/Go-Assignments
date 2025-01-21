package main

import (
	"fmt"
	"sync"
)

func processConversation(conversation string, aliceConversationChan chan<- string, bobConversationChan chan<- string, stopChan chan<- bool) {
	message := ""
	for _, ch := range conversation {
		if ch == '$' {
			aliceConversationChan <- message
			message = ""
		} else if ch == '#' {
			bobConversationChan <- message
			message = ""
		} else if ch == '^' {
			stopChan <- true
			return
		} else {
			message += string(ch)
		}
	}

	stopChan <- true
}

func printConversation(aliceConversationChan <-chan string, bobConversationChan <-chan string, stopChan <-chan bool, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case message := <-aliceConversationChan:
			fmt.Printf("alice : %s,", message)
		case message := <-bobConversationChan:
			fmt.Printf("bob : %s,", message)
		case <-stopChan:
			fmt.Printf("\033[1D \n")
			return
		}
	}
}

func main() {
	conversation := "helloBob$helloalice#howareyou?#Iamgood.howareyou?$"

	aliceConversationChan := make(chan string)
	bobConversationChan := make(chan string)
	stopChan := make(chan bool)
	var wg sync.WaitGroup

	go processConversation(conversation, aliceConversationChan, bobConversationChan, stopChan)

	wg.Add(1)

	go printConversation(aliceConversationChan, bobConversationChan, stopChan, &wg)

	wg.Wait()
}
