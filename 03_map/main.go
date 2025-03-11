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
}

func printMap(m map[string]string) {
	for k, v := range m {
		fmt.Println(k, v)
	}
}
