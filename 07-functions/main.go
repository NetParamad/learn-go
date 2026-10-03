package main

import "fmt"

// 07 - functions โน้ต: func ระบุ type รับ/คืน, multiple return ใช้กับ error
func add(a, b int) int { return a + b }

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("div by zero หารด้วยศูนย์ไม่ได้")
	}
	return a / b, nil
}

func main() {
	fmt.Println("add:", add(2, 3))

	v, err := divide(10, 2)
	fmt.Println("div:", v, "err:", err)

	_, err = divide(1, 0)
	fmt.Println("err case:", err)
}
