/*
This snippet demonstrates how we can declare a struct to unmarshal a
JSON to it and only extrac the fields we're actually interested in.

This technique uses interfaces which are equal to any for the fields
we're not interested in.

*/

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type postsResp struct {
    UserId any
    Id any // We don't care about this field
    Title string `json:"title"` // Struct tag which tells that the key in the document is title
    Body any // We don't care about this field
}

var posts []postsResp

func perror(e error)  {
    if e != nil {
        log.Fatal(e)
    }
}

func main()  {
    url := "https://jsonplaceholder.typicode.com/posts"
    resp, err := http.Get(url)
    perror(err)
    defer resp.Body.Close()
    body, err := io.ReadAll(resp.Body)
    perror(err)

    err = json.Unmarshal(body, &posts)
    perror(err)
    fmt.Println("Posts:")
    for _, post := range posts {
        fmt.Println(post.Title)
        fmt.Println(post.Id)
    }
}
