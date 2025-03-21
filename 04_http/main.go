package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct {}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}

    lw := logWriter{}

	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (n int, err error){
    fmt.Println("Congrats, you just implemented the Writer interface!")
    return 1, nil
}
