package exercise15

import (
	"fmt"
	"time"
)

func Run() {
	var nonChannel string = "eagle"
	fmt.Printf("\nnonChannel: %v\n", nonChannel)

	channel1 := make(chan string)
	channel2 := make(chan string)

	// <-chan to receive from channel, chain<- to send to channel. can determine this for the channel within the maker
	// ie channel3 := make(chan<- string) only allows receiving from channel (directional channels vs bidirectional channels)

	go func() {
		time.Sleep(time.Duration(2 * time.Second))
		channel1 <- "crow"
	}()

	fmt.Printf("\nchannel1: %v\n", <-channel1)

	go func() {
		time.Sleep(time.Duration(time.Second))
		channel2 <- "chicken"
	}()

	fmt.Printf("\nchannel2: %v\n", <-channel2)
}
