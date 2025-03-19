package main

import ("slices";"fmt")

func main() {
    s := []int{1, 2, 3, 4}

    // Delete from slice
    fmt.Printf("Original slice: %v\n", s)
    s = slices.Delete(s, 1, 2)
    fmt.Printf("Modified slice: %v\n", s)

    // Check if element in slice
    // https://pkg.go.dev/golang.org/x/exp/slices#Contains
    if slices.Contains(s, 1) {
        fmt.Println("1 in s")
    }
}
