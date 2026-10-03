package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

// 16 - todo CLI โน้ต: CRUD ผ่าน flag + เก็บไฟล์ (แทน localStorage)
type Todo struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
}

const file = "todos.json"

func load() []Todo {
	data, err := os.ReadFile(file)
	if err != nil {
		return nil
	}
	var t []Todo
	_ = json.Unmarshal(data, &t)
	return t
}

func save(t []Todo) {
	data, _ := json.MarshalIndent(t, "", "  ")
	_ = os.WriteFile(file, data, 0644)
}

func main() {
	add := flag.String("add", "", "เพิ่ม todo เช่น -add=\"อ่าน Go\"")
	list := flag.Bool("list", false, "แสดงทั้งหมด")
	flag.Parse()

	todos := load()
	if *add != "" {
		todos = append(todos, Todo{ID: len(todos) + 1, Text: *add})
		save(todos)
		fmt.Println("added:", *add)
		return
	}
	if *list {
		for _, t := range todos {
			fmt.Printf("%d. %s done=%v\n", t.ID, t.Text, t.Done)
		}
		return
	}
	fmt.Println("ใช้: go run . -add=\"...\" | -list")
}
