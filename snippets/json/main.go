package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	type ColorGroup struct {
		ID     int
		Name   string
		Colors []string
		Test string 
	}

	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
        Test: "test",
    }

	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)
}

