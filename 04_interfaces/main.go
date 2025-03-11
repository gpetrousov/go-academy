// A better example

/*
https://go.dev/tour/methods/9

package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type Vertex struct {
	X, Y float64
}

type MyFloat float64


func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}
	
	fmt.Printf("f=%v, v=%v, a=%v\n", f, v, a)

	a = f  // a MyFloat implements Abser
	fmt.Printf("a can implement f; a=%v\n", a)
	
	fmt.Println(a.Abs())
	
	a = &v // a *Vertex implements Abser
	fmt.Printf("a can implement v; a=%v\n", a)
	
	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	//a = v

	fmt.Println(a.Abs())
}

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
*/



package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

}

func (englishBot) getGreeting() string {
	return "Hello"
}

func (spanishBot) getGreeting() string {
	return "Holla"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
