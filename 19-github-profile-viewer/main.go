package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"time"
)

// 19 - github profile viewer โน้ต: api.github.com/users/:name
type Profile struct {
	Login string `json:"login"`
	Name  string `json:"name"`
	Bio   string `json:"bio"`
	Repos int    `json:"public_repos"`
}

func main() {
	user := flag.String("user", "golang", "github username")
	flag.Parse()

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get("https://api.github.com/users/" + *user)
	if err != nil {
		fmt.Println("fetch err (ต้องมีเน็ต):", err)
		return
	}
	defer resp.Body.Close()

	var p Profile
	if err := json.NewDecoder(resp.Body).Decode(&p); err != nil {
		fmt.Println("decode err:", err)
		return
	}
	fmt.Printf("%s (%s) repos=%d bio=%s\n", p.Login, p.Name, p.Repos, p.Bio)
}
