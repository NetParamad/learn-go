package main

import "fmt"

// 03 - data types โน้ต: static type, zero value ต่างจาก JS
func main() {
	var i int // zero value = 0
	var f float64 = 19.99
	var s string // zero value = ""
	var b bool   // zero value = false

	fmt.Printf("int:%d float:%.2f string:%q bool:%t\n", i, f, s, b)

	// type inference ด้วย :=
	name := "Go"
	n := 42
	fmt.Printf("%s %d\n", name, n)
}
