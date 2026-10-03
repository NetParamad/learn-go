package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// 14 - files & json โน้ต: os + encoding/json แทน localStorage
type Todo struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
}

func main() {
	todos := []Todo{{1, "เรียน Go", false}, {2, "เขียน CLI", true}}

	data, _ := json.MarshalIndent(todos, "", "  ")
	_ = os.WriteFile("todos.json", data, 0644)
	fmt.Println("wrote todos.json")

	read, _ := os.ReadFile("todos.json")
	var loaded []Todo
	_ = json.Unmarshal(read, &loaded)
	fmt.Println("loaded:", loaded)
}
