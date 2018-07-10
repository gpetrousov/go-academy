package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	c := make(chan string)
	links := []string{
		"http://facebook.com",
		"http://google.com",
		"http://golang.org",
	}

	for _, link := range links {
		go checkLink(link, c)
	}
	for {
		go checkLink(<-c, c)
	}
}

func checkLink(link string, c chan string) {
	time.Sleep(5 * time.Second)
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "tango down")
		c <- link
		return
	}
	fmt.Println(link, "is up")
	c <- link
}
