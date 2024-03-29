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

	c := make(chan string) // create new channel

	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c { // infinite loop; only loops anytime we receive something through a channel
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l) // need () at end to invoke function literal
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}
	fmt.Println(link, "is up!")
	c <- link
}
