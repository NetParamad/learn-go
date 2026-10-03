package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"strings"
	"time"
)

// 17 - product search โน้ต: ดึง API + filter ฝั่ง client
type Product struct {
	ID    int     `json:"id"`
	Title string  `json:"title"`
	Price float64 `json:"price"`
}

func main() {
	q := flag.String("q", "", "คำค้น เช่น -q=\"shirt\"")
	flag.Parse()

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get("https://fakestoreapi.com/products")
	if err != nil {
		fmt.Println("fetch err (ต้องมีเน็ต):", err)
		return
	}
	defer resp.Body.Close()

	var items []Product
	if err := json.NewDecoder(resp.Body).Decode(&items); err != nil {
		fmt.Println("decode err:", err)
		return
	}

	for _, p := range items {
		if *q == "" || strings.Contains(strings.ToLower(p.Title), strings.ToLower(*q)) {
			fmt.Printf("%d %s %.2f\n", p.ID, p.Title, p.Price)
		}
	}
}
