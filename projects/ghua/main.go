/*
Output: count the actions per repository

eventTYpe   repository  action  counter
map[string]any

*/

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type GitHubRepo struct {
    Url string `json:"url"`
}

type GitHubPayload struct {
    Action string `json:"action"`
}

type GitHubEvent struct {
    Type string `json:"type"`
    Repo GitHubRepo `json:"repo"`
    Payload GitHubPayload `json:"payload"`
}

type UserActions struct {
    repository string
    action  string
    counter int
}

func perror(e error)  {
    if e != nil {
        log.Fatal(e)
    }
}

func main()  {

    // Fetch data
    uName := os.Args[1]
    url := fmt.Sprintf("https://api.github.com/users/%v/events/public", uName)
    resp, err := http.Get(url)
    perror(err)

    // Extract body
    defer resp.Body.Close()
    body, err := io.ReadAll(resp.Body)
    perror(err)

    // Unmarshal events
    var ghEvents []GitHubEvent 
    err = json.Unmarshal(body, &ghEvents)
    perror(err)

    userStats := make(map[string][]UserActions)
    for _, e := range ghEvents {
        counterIncreased := false
        // fmt.Printf("Event Type: %s\n", e.Type)
        // fmt.Printf("Event Repo: %s\n", e.Repo.Url)
        // fmt.Printf("Event Action: %s\n", e.Payload.Action)
        // fmt.Println("---")
        /*
        add_found = False

        if eventType in map
            if repoUrl in slice AND action in slice
                add_found = True
                mark_counter()
            if event_found:
                add event to slice
                makr
        else
            Add new eventType object to map
        */
        _, ok := userStats[e.Type]
        if ok {
            for i, itm := range userStats[e.Type] {
                if (e.Repo.Url == itm.repository) && (e.Payload.Action == itm.action) {
                    userStats[e.Type][i].counter += 1
                    counterIncreased = true
                    break
                }
            }
            if !counterIncreased {
                userStats[e.Type] = append(userStats[e.Type], UserActions{repository: e.Repo.Url, action: e.Payload.Action, counter: 1})
            }
        } else {
            // Event not in map
            newUserActions := []UserActions{
                {
                    repository: e.Repo.Url,
                    action: e.Payload.Action,
                    counter: 1,
                },
            }
            userStats[e.Type] = newUserActions
            }
        }
        fmt.Println(userStats)
    }
