package main

import "fmt"

// 09 - maps & structs โน้ต: map แทน object, struct คือ type มี field
type User struct {
	ID   int
	Name string
}

func main() {
	// map แทน object ของ JS
	scores := map[string]int{"Go": 90, "JS": 80}
	scores["TS"] = 85
	fmt.Println(scores, "Go:", scores["Go"])

	// struct
	u := User{ID: 1, Name: "Net"}
	fmt.Println(u)

	// slice of structs
	users := []User{{1, "A"}, {2, "B"}}
	fmt.Println(users)
}
