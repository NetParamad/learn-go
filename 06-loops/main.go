package main

import "fmt"

// 06 - loops โน้ต: มีแค่ for, ใช้แทน while ด้วย
func main() {
	// classic for
	for i := 0; i < 3; i++ {
		fmt.Println("i:", i)
	}

	// while-style
	n := 3
	for n > 0 {
		fmt.Println("n:", n)
		n--
	}

	// range
	langs := []string{"Go", "JS", "TS"}
	for idx, v := range langs {
		fmt.Println(idx, v)
	}
}
