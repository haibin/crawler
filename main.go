package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type Lang struct {
	Name string
	Year int
	URL  string
}

var langs = []Lang{
	{"Python", 1991, "http://python.org/"},
	{"Ruby", 1995, "http://www.ruby-lang.org/en/"},
	{"Scala", 2003, "http://www.scala-lang.org/"},
	{"GO", 2009, "http://golang.org/"},
}

func main() {
	start := time.Now()

	c := make(chan string)
	n := 0
	do(func(lang Lang) {
		n++
		go count(lang.Name, lang.URL, c)
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

func do(f func(Lang)) {
	for _, l := range langs {
		f(l)
	}
}

func count(name, url string, c chan<- string) {
	start := time.Now()
	r, err := http.Get(url)
	if err != nil {
		c <- fmt.Sprintf("%s: %s", name, err)
		return
	}
	n, _ := io.Copy(ioutil.Discard, r.Body)
	r.Body.Close()
	c <- fmt.Sprintf("%s %d [%.2fs]\n", name, n, time.Since(start).Seconds())
}
