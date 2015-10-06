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

type Crawler interface {
	Do(c chan<- string)
}

func New(langs []Lang) Crawler {
	return &simpleCrawler{langs}
}

type simpleCrawler struct {
	langs []Lang
}

func (cl *simpleCrawler) Do(c chan<- string) {
	for _, l := range cl.langs {
		go fetch(l.Name, l.URL, c)
	}
}

func fetch(name, url string, c chan<- string) {
	start := time.Now()
	res, err := http.Get(url)
	if err != nil {
		c <- fmt.Sprintf("%s: %s", name, err)
		return
	}
	// res.Body is an io.ReadCloser
	// Discard is an io.Writer
	n, _ := io.Copy(ioutil.Discard, res.Body)
	res.Body.Close()
	c <- fmt.Sprintf("%s %d [%.2fs]\n", name, n, time.Since(start).Seconds())
}
