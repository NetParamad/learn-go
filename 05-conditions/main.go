package main

import "fmt"

// 05 - conditions โน้ต: if ต้องมี bool, switch ไม่ต้อง break
func main() {
	age := 20
	if age >= 18 {
		fmt.Println("adult: ผู้ใหญ่")
	} else {
		fmt.Println("minor: เยาวชน")
	}

	// if + short statement
	if n := 7; n%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}

	// switch
	day := "Sat"
	switch day {
	case "Sat", "Sun":
		fmt.Println("weekend")
	default:
		fmt.Println("weekday")
	}
}
