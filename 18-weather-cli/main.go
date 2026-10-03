package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// 18 - weather CLI โน้ต: open-meteo ไม่ต้องใช้ key (แทน weather-app)
func main() {
	city := flag.String("city", "Bangkok", "ชื่อเมือง (demo ใช้ lat/lon Bangkok)")
	flag.Parse()

	// demo: Bangkok lat/lon (geocoding จริงค่อยต่อยอด)
	_ = city
	api := "https://api.open-meteo.com/v1/forecast?latitude=13.75&longitude=100.5&current=temperature_2m,weather_code"

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(api)
	if err != nil {
		fmt.Println("fetch err (ต้องมีเน็ต):", err)
		return
	}
	defer resp.Body.Close()

	var out map[string]any
	_ = json.NewDecoder(resp.Body).Decode(&out)
	cur, _ := out["current"].(map[string]any)
	fmt.Println("query:", url.QueryEscape(*city))
	fmt.Println("current:", cur)
}
