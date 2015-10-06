package main

import (
	"fmt"
	"time"

	"github.com/haibin/crawler"
)

func main() {
	start := time.Now()

	c := make(chan string)
	n := 0
	crawler.Do(func(lang crawler.Lang) {
		n++
		go crawler.Count(lang.Name, lang.URL, c)
	})

	// time.After creates a channel
	timeout := time.After(2 * time.Second)
	for i := 0; i < n; i++ {
		select {
		case result := <-c:
			fmt.Print(result)
		case <-timeout:
			fmt.Println("Time out")
			return
		}
	}

	fmt.Printf("%.2fs total\n", time.Since(start).Seconds())
}
