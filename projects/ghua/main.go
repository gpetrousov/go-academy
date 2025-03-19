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
	"strings"
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
        fmt.Printf("User: %s\n", uName)
        for ghEvent, uData := range userStats{
            switch ghEvent {
            case "CreateEvent":
                for _, v := range uData{
                    fmt.Printf("- Created %d Branches in %s\n", v.counter, v.repository)
                }
            case "IssuesEvent":
                for _, v := range uData{
                    fmt.Printf("- %s Issues in %s %d times\n", strings.Title(v.action), v.repository, v.counter)
                }
            case "WatchEvent":
                for _, v := range uData{
                    fmt.Printf("- %s watching %s %d times\n", strings.Title(v.action), v.repository, v.counter)
                }
            case "PullRequestEvent":
                for _, v := range uData{
                    fmt.Printf("- %s %d PRs in %s\n", strings.Title(v.action), v.counter, v.repository)
                }
            case "IssueCommentEvent":
                for _, v := range uData{
                    fmt.Printf("- Posted %d Comment(s) in %s\n", v.counter, v.repository)
                }
            case "DeleteEvent":
                for _, v := range uData{
                    fmt.Printf("- Deleted %d times branches in %s\n", v.counter, v.repository)
                }
            case "PushEvent":
                for _, v := range uData{
                    fmt.Printf("- PUshed %d times to %s\n", v.counter, v.repository)
                }
            }
        }
    }
