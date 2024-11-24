package main

import (
	"fmt"
	"sync"
	"time"
)

// Create a struct of type Message
type Message struct {
	chats   []string
	friends []string
}

// Function to get user name
func getUserByName(name string) string {
	time.Sleep(time.Second * 1)

	return fmt.Sprintf("userName printed after 1 second: %s ", name)
}

// Function to get user chats
func getUserChats(userName string, ch chan<- *Message, wg *sync.WaitGroup) {
	//wg.Add(1) //Adds to WaitGroup
	defer wg.Done() // Invokes Done message by the channel after function is exectuted
	time.Sleep(time.Second * 2)

	ch <- &Message{
		chats: []string{
			"Hey! I'm awesome",
			"I love singing",
			"I love dancing",
		},
	}

	//wg.Done() //Done message by the channel for the WaitGroup
}

// Funtion to get user friends
func getUserFriends(userName string, ch chan<- *Message, wg *sync.WaitGroup) {
	//wg.Add(1) //Adds to WaitGroup
	defer wg.Done() // Invokes Done message by the channel after function is exectuted
	time.Sleep(time.Second * 3)

	ch <- &Message{
		friends: []string{
			"John",
			"Jane",
			"Joe",
		},
	}

	//wg.Done() //Done message by the channel for the WaitGroup
}
