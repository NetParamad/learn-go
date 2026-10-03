package main

import (
	"fmt"
	"sync"
)

// 13 - goroutines & channels โน้ต: go f() รัน concurrently, channel ส่งค่าระหว่างกัน
func worker(id int, ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- fmt.Sprintf("งาน %d เสร็จ", id)
}

func main() {
	ch := make(chan string, 3)
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, ch, &wg)
	}
	wg.Wait()
	close(ch)

	for msg := range ch {
		fmt.Println(msg)
	}
}
