/*
Output

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

type GitHubEvent struct {
    Type string `json:"type"`
    Repo json.RawMessage `json:"repo"`
    Payload json.RawMessage `json:"payload"`
}

type GitHubRepo struct {
    Url string `json:"url"`
}

type GitHubPayload struct {
    Action string `json:"action"`
}

type UserActions struct {
    repository string
    action  string
    counter int
}

type PublicUserEvents map[string][]UserActions

func perror(e error)  {
    if e != nil {
        log.Fatal(e)
    }
}

func (m PublicUserEvents)contains(s string) bool {
    _, ok := m[s]
    if ok {
        return true
    } else {
        return false
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

    var ghRepo GitHubRepo
    var ghPayload GitHubPayload
    publicUserEvents := make(PublicUserEvents)
    for _, e := range ghEvents {
        err = json.Unmarshal(e.Repo, &ghRepo)
        perror(err)
        repoUrl := ghRepo.Url
        err = json.Unmarshal(e.Payload, &ghPayload)
        perror(err)
        action := ghPayload.Action
        eventType := e.Type
        matched := false

        // If type is in publicUserEvents
        // Check if repository matches
        // Check if action matches
        // If all match => increase counter

        // If one does not match => add UserActions struct

        if publicUserEvents.contains(eventType) {
            // Check repo match
            for _, ua := range publicUserEvents[eventType] {
                if ua.repository == repoUrl && ua.action == action {
                    ua.counter += 1
                    matched = true
                    break
                }
            }
            if !matched {
                newUserAction := UserActions {
                    repository: repoUrl,
                    action: action,
                    counter: 0,
                }
                publicUserEvents[eventType] = append(publicUserEvents[eventType], newUserAction)
                matched = false
            }
        } else {
            newUserAction := UserActions {
                repository: repoUrl,
                action: action,
                counter: 0,
            }
            publicUserEvents[eventType] = append(publicUserEvents[eventType], newUserAction)
        }
    }
    fmt.Println(publicUserEvents["PushEvent"])
}
