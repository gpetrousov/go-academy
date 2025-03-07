package main

import "fmt"

var c, python bool = true, false
var java = "no!"

// Type casting
var q int = 54
var w float64 = float64(q) + 0.0001

// Type inference
var ii int

// Constant declaration
const Pi = 3.142

const (
    Big = 1 << 3
    // Big = 1 << 300 overflows
)

func add(x int, y int) int {
    return x + y
}

func add_another(x, y int) int {
    return x + y
}

func swap(x, y string) (string, string) {
    return y, x
}

func naked_return(sum int) (x, y int) {
    x = sum/2
    y = sum*2
    return
}

func needsInt(x int) int {
    return x*10
}

func main() {
    var i int

    /*
    Implicit type declaration
    Only inside function
    */ 
    c := 10

    // Type inference
    jj := ii
    jjj := 3.142

    fmt.Println("Hello world!")
    fmt.Println(add(4, 12))
    fmt.Println(add_another(4, 12))
    fmt.Println(swap("hello", "world"))
    fmt.Println(naked_return(48))

    // Print formatting
    fmt.Printf("%v %v %v %q\n", i, c, python, java)

    // Casted values
    fmt.Printf("%v, %v\n", q, w)

    // Type inference
    fmt.Printf("jj type: %T\njjj type: %T\n", jj, jjj)

    // Constant usage
    fmt.Println(Pi)
    fmt.Println(Big)
    fmt.Println(needsInt(Big))
}
