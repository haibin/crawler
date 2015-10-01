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

var langs []Lang

func main() {
	start := time.Now()

	langs = []Lang{
		{"Python", 1991, "http://python.org/"},
		{"Ruby", 1995, "http://www.ruby-lang.org/en/"},
		{"Scala", 2003, "http://www.scala-lang.org/"},
		{"GO", 2009, "http://golang.org/"},
	}

	do(func(lang Lang) {
		count(lang.Name, lang.URL)
	})
	fmt.Printf("%.2fs total\n", time.Since(start).Seconds())
}

func do(f func(Lang)) {
	for _, l := range langs {
		f(l)
	}
}

func count(name, url string) {
	start := time.Now()
	r, err := http.Get(url)
	if err != nil {
		fmt.Printf("%s: %s", name, err)
		return
	}
	n, _ := io.Copy(ioutil.Discard, r.Body)
	r.Body.Close()
	fmt.Printf("%s %d [%.2fs]\n", name, n, time.Since(start).Seconds())
}
