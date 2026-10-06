package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// 18 - weather CLI โน้ต: geocoding + open-meteo ไม่ต้องใช้ key (แทน weather-app)
func main() {
	city := flag.String("city", "Bangkok", "ชื่อเมือง เช่น -city=\"Chiang Mai\"")
	flag.Parse()

	client := &http.Client{Timeout: 10 * time.Second}

	// 1) geocoding: ชื่อเมือง -> lat/lon
	geoAPI := "https://geocoding-api.open-meteo.com/v1/search?name=" +
		url.QueryEscape(*city) + "&count=1&language=en&format=json"
	geoResp, err := client.Get(geoAPI)
	if err != nil {
		fmt.Println("geocode fetch err (ต้องมีเน็ต):", err)
		return
	}
	defer geoResp.Body.Close()
	if geoResp.StatusCode != http.StatusOK {
		fmt.Println("geocode http err:", geoResp.Status)
		return
	}
	var geo struct {
		Results []struct {
			Name      string  `json:"name"`
			Country   string  `json:"country"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"results"`
	}
	if err := json.NewDecoder(geoResp.Body).Decode(&geo); err != nil {
		fmt.Println("geocode decode err:", err)
		return
	}
	if len(geo.Results) == 0 {
		fmt.Println("ไม่พบเมือง:", *city)
		return
	}
	place := geo.Results[0]

	// 2) forecast: lat/lon -> อากาศปัจจุบัน
	api := fmt.Sprintf(
		"https://api.open-meteo.com/v1/forecast?latitude=%v&longitude=%v&current=temperature_2m,weather_code",
		place.Latitude, place.Longitude)
	resp, err := client.Get(api)
	if err != nil {
		fmt.Println("fetch err (ต้องมีเน็ต):", err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("weather http err:", resp.Status)
		return
	}

	var out map[string]any
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		fmt.Println("decode err:", err)
		return
	}
	cur, ok := out["current"].(map[string]any)
	if !ok {
		fmt.Println("decode err: ไม่มี field current")
		return
	}
	fmt.Printf("query: %s (%s)\n", place.Name, place.Country)
	fmt.Println("current:", cur)
}
