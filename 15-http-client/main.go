package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// 15 - http client โน้ต: net/http แทน fetch
type Product struct {
	ID    int     `json:"id"`
	Title string  `json:"title"`
	Price float64 `json:"price"`
}

func main() {
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get("https://fakestoreapi.com/products/1")
	if err != nil {
		fmt.Println("fetch err (ต้องมีเน็ต):", err)
		return
	}
	defer resp.Body.Close()

	var p Product
	if err := json.NewDecoder(resp.Body).Decode(&p); err != nil {
		fmt.Println("decode err:", err)
		return
	}
	fmt.Printf("product: %+v\n", p)
}
