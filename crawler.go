package crawler

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type Lang struct {
	Name string
	URL  string
}

var langs = []Lang{
	{"Python", "http://python.org/"},
	{"Ruby", "http://www.ruby-lang.org/en/"},
	{"Scala", "http://www.scala-lang.org/"},
	{"GO", "http://golang.org/"},
}

func Do(f func(Lang)) {
	for _, l := range langs {
		f(l)
	}
}

func Count(name, url string, c chan<- string) {
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
