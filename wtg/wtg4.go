package main

import "fmt"
import "math"

type Simpler interface {
	Get() int
	Set(int)
}

type Simple struct {
	i int
	Simpler
}

func (s Simple) Get() int {
	return s.i
}

func (s *Simple) Set(x int) {
	s.i = x
}

type Square struct {
	side float32
}

type Circle struct {
	radius float32
}

type Shaper interface {
	Area() float32
}

type PeriInteface interface {
	Perimeter() float32
}

type Triangle struct {
	base   float32
	height float32
	Shaper
}

func main() {
	var a Simpler = &Simple{i: 42}
	fmt.Println(a.Get())
	a.Set(23)
	fmt.Println(a.Get())
	var areaIntf Shaper
	sq1 := new(Square)
	sq1.side = 5
	areaIntf = sq1
	if t, ok := areaIntf.(*Square); ok {
		fmt.Printf("The type of areaIntf is: %T\n", t)
	}
	if u, ok := areaIntf.(*Circle); ok {
		fmt.Printf("The type of areaIntf is: %T\n", u)
	} else {
		fmt.Printf("areaIntf does not contain a variable of type Circle\n")
	}
	switch t := areaIntf.(type) {
	case *Square:
		fmt.Printf("Type Square %T with Value %v\n", t, t)
	case *Circle:
		fmt.Printf("Type Circle %T with Value %v\n", t, t)
	case nil:
		fmt.Printf("nil value: nothing to check?\n")
	default:
		fmt.Printf("Unexpected type %T\n", t)
	}
	t := &Triangle{base: 4.2, height: 3.0}
	fmt.Println(t.Area())
	p := &Square{3.2}
	fmt.Println(p.Perimeter())

}

func (sq *Square) Area() float32 {
	return sq.side * sq.side

}
func (sq *Square) Perimeter() float32 {
	return 4 * sq.side
}

func (c *Circle) Area() float32 {
	return 2 * c.radius * math.Pi
}

func (t *Triangle) Area() float32 {
	return 0.5 * t.base * t.height
}
