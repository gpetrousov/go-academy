package main

import "fmt"

func main() {
	// var colors map[string]string

    //colors := make(map[string]string)

	colors := map[string]string{
		"red":   "#AAAAAA",
		"green": "#FFFFFF",
	}

    // Add item to map
	colors["blue"] = "#123456"

	printMap(colors)


    /*
    Check if key in map - Use the comma, ok syntax
    https://golang.cafe/blog/how-to-check-if-a-map-contains-a-key-in-go
    */

    color, ok := colors["red"]
    if ok {
        fmt.Printf("%s in list\n", color)
    } else {
        fmt.Printf("%s not in list\n", color)
    }

    color, ok = colors["cyan"]
    if ok {
        fmt.Printf("%s in list\n", color)
    } else {
        fmt.Printf("%s not in list\n", "cyan")
    }
}

func printMap(m map[string]string) {
	for k, v := range m {
		fmt.Println(k, v)
	}
}
