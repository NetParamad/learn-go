package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

// 20 - http server โน้ต: net/http server แทน DOM/events
func main() {
	port := flag.String("port", "", "พอร์ต (default :8080 หรือ env PORT)")
	flag.Parse()
	addr := ":" + firstNonEmpty(*port, os.Getenv("PORT"), "8080")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(map[string]string{"msg": "สวัสดีจาก Go server"}); err != nil {
			log.Println("encode err:", err)
		}
	})
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "ok")
	})

	fmt.Println("listen", addr, "-> curl localhost"+addr+"/")
	log.Fatal(http.ListenAndServe(addr, nil))
}

func firstNonEmpty(vals ...string) string {
	for _, v := range vals {
		if v != "" {
			return v
		}
	}
	return ""
}
