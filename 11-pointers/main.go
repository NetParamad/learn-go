package main

import "fmt"

// 11 - pointers โน้ต: *T คือ pointer, & เอาที่อยู่, * ดึงค่า
func inc(n *int) { *n++ } // แก้ค่าต้นฉบับได้

func main() {
	x := 10
	fmt.Println("before:", x)
	inc(&x)
	fmt.Println("after:", x)

	p := &x
	fmt.Println("ptr:", p, "value:", *p)
}
