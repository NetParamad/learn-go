package main

import (
	"errors"
	"fmt"
)

// 12 - errors โน้ต: error คือค่า return ปกติ, เช็คด้วย if err != nil
func age(n int) (string, error) {
	if n < 0 {
		return "", errors.New("age ติดลบไม่ได้")
	}
	if n < 18 {
		return "minor", nil
	}
	return "adult", nil
}

func main() {
	for _, n := range []int{20, 10, -1} {
		s, err := age(n)
		if err != nil {
			fmt.Println(n, "err:", err)
			continue
		}
		fmt.Println(n, s)
	}
}
