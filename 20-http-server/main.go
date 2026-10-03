package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// 20 - http server โน้ต: net/http server แทน DOM/events
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]string{"msg": "สวัสดีจาก Go server"})
	})
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "ok")
	})

	fmt.Println("listen :8080 -> curl localhost:8080/")
	_ = http.ListenAndServe(":8080", nil)
}
