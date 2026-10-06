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
		if !os.IsNotExist(err) {
			fmt.Println("load err:", err)
		}
		return nil
	}
	var t []Todo
	if err := json.Unmarshal(data, &t); err != nil {
		fmt.Println("parse err (ไฟล์เสีย เริ่มใหม่):", err)
		return nil
	}
	return t
}

func save(t []Todo) error {
	data, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(file, data, 0644)
}

func nextID(t []Todo) int {
	max := 0
	for _, v := range t {
		if v.ID > max {
			max = v.ID
		}
	}
	return max + 1
}

func main() {
	add := flag.String("add", "", "เพิ่ม todo เช่น -add=\"อ่าน Go\"")
	list := flag.Bool("list", false, "แสดงทั้งหมด")
	flag.Parse()

	todos := load()
	if *add != "" {
		todos = append(todos, Todo{ID: nextID(todos), Text: *add})
		if err := save(todos); err != nil {
			fmt.Println("save err:", err)
			return
		}
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
