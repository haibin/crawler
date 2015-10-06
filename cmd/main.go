package main

import (
	"fmt"
	"time"

	"github.com/haibin/crawler"
)

func main() {
	start := time.Now()

	var langs = []crawler.Lang{
		{"Python", "http://python.org/"},
		{"Ruby", "http://www.ruby-lang.org/en/"},
		{"Scala", "http://www.scala-lang.org/"},
		{"GO", "http://golang.org/"},
	}

	crawler := crawler.New(langs)

	c := make(chan string)
	crawler.Crawl(c)

	// time.After creates a channel
	timeout := time.After(3 * time.Second)
	for i := 0; i < len(langs); i++ {
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
