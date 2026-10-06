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
	todos := []Todo{{ID: 1, Text: "เรียน Go", Done: false}, {ID: 2, Text: "เขียน CLI", Done: true}}

	data, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		fmt.Println("marshal err:", err)
		return
	}
	if err := os.WriteFile("todos.json", data, 0644); err != nil {
		fmt.Println("write err:", err)
		return
	}
	fmt.Println("wrote todos.json")

	read, err := os.ReadFile("todos.json")
	if err != nil {
		fmt.Println("read err:", err)
		return
	}
	var loaded []Todo
	if err := json.Unmarshal(read, &loaded); err != nil {
		fmt.Println("unmarshal err:", err)
		return
	}
	fmt.Println("loaded:", loaded)
}
