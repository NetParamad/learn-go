package main

import "fmt"

// 10 - methods & interfaces โน้ต: method ผูกกับ type, interface คือชุด method
type Rect struct{ W, H float64 }

func (r Rect) Area() float64 { return r.W * r.H }

type Shaper interface{ Area() float64 }

func printArea(s Shaper) { fmt.Println("area:", s.Area()) }

func main() {
	r := Rect{W: 3, H: 4}
	fmt.Println("direct:", r.Area())
	printArea(r) // ใช้ผ่าน interface ได้
}
