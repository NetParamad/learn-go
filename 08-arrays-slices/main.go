package main

import "fmt"

// 08 - arrays & slices โน้ต: array ขนาด fix, slice ยืดได้ (ใช้บ่อยกว่า)
func main() {
	var arr [3]int = [3]int{1, 2, 3} // array: len ตายตัว
	fmt.Println("arr:", arr, "len:", len(arr))

	s := []int{1, 2, 3} // slice: ยืดได้
	s = append(s, 4, 5)
	fmt.Println("slice:", s, "len:", len(s), "cap:", cap(s))

	// slice expression
	fmt.Println("s[1:3]:", s[1:3])

	// make
	m := make([]int, 0, 2)
	fmt.Println("make:", m, len(m), cap(m))
}
