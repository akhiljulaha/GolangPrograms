package main

import (
	"fmt"
	"time"
)

/* Step 5
In an unbuffered channel, both sending and receiving are blocking operations.
To avoid this we can use a buffered channel.

Defiination: A buffered channel in Go is a channel with a fixed capacity that allows you to send values without immediately blocking, as long as the buffer isn’t full.
*/



func emailSender(emailChan chan string, done chan bool){
	defer func ()  {done <- true}() 
	for email := range emailChan{
		fmt.Println("sending email to", email)
		time.Sleep(time.Second)
	}
	// done <- true
}

func main() {


	emailChan := make(chan string, 100)  // buffer capacity

	done := make(chan bool)

	go emailSender(emailChan, done)
	
	for i:=0; i<5; i++{
		emailChan <- fmt.Sprintf("%d@gmail.com", i)
}
fmt.Println("done sending")

close(emailChan) // Important to close the channel because range will run infinite times, so if we will not close so will get the deadlock error
<-done

}

 