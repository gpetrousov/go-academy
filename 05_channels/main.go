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

    for l := range c {
        go func(link string) {
            time.Sleep(3 *time.Second)
            checkLink(link, c)
        }(l)
    }
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
        fmt.Printf("%s => tango down\n", link)
        c <- link
		return
	}
	// fmt.Println(link, "is up")
    fmt.Printf("%s => is alive\n", link)
    c <- link
}
