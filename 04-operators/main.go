package main

import "fmt"

// 04 - operators โน้ต: เหมือน JS แต่ไม่มี ++x แบบ expression, % ใช้กับ int
func main() {
	a, b := 10, 3
	fmt.Println("add:", a+b)
	fmt.Println("div:", a/b) // int division = 3
	fmt.Println("mod:", a%b)

	// comparison + logical เหมือน JS
	fmt.Println("eq:", a == b, "gt:", a > b)
	fmt.Println("and:", a > 5 && b < 5)
}
