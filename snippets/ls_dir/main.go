package main

import (
	"fmt"
	"log"
	"os"
)

func main()  {
    jsonFile := "godos.json"
    b, err := os.ReadFile(jsonFile)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Length: %v\n", len(b))
}
