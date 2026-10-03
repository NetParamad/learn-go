package main

import "fmt"

// 02 - variables โน้ต: := ประกาศสั้น, var ระบุ type, const เปลี่ยนไม่ได้
func main() {
	name := "Net"  // string, เปลี่ยนได้
	const age = 25 // const ห้าม assign ใหม่
	var city string = "Bangkok"

	name = "Paramad" // ok
	// age = 26 // Error: cannot assign to age

	fmt.Println("name:", name)
	fmt.Println("age:", age)
	fmt.Println("city:", city)
}
