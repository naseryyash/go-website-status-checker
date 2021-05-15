package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	// make a channel of type string
	c := make(chan string)

	for _, link := range links {
		// spawn a new Go-Routine
		go checkLink(link, c)
	}

	for l := range c {
		go func(link string) {
			// sleep routine for a while
			// introduce a delay between 2 status checks
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, " might be down!!")
		// send message into the channel
		c <- link
		return
	}

	fmt.Println(link, " is up")
	// send message into the channel
	c <- link
}
