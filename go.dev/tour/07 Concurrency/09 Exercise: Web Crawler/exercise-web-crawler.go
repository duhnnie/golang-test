package main

import (
	"fmt"
	"sync"
	"time"
)

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

type fakeResult struct {
	body string
	urls []string
}

type fakeFetcher map[string]*fakeResult

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}

	return "", nil, fmt.Errorf("not found: %s", url)
}

var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}

type cache struct {
	m sync.Mutex
	v map[string]int
}

var myCache = cache{ v: make(map[string]int) }

func (c *cache) Inc(url string) {
	c.m.Lock()
	c.v[url]++
	c.m.Unlock()
}

func (c *cache) IsCrawled(url string) bool {
	c.m.Lock()
	defer c.m.Unlock()
	_, ok := c.v[url]
	return ok
}

func Crawl(url string, depth int, fetcher Fetcher) {
	if myCache.IsCrawled(url) {
		return
	}

	myCache.Inc(url)

	if depth <= 0 {
		return
	}

	body, urls, err := fetcher.Fetch(url)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Found: %s %q\n", url, body)

	for _, url := range urls {
		go Crawl(url, depth - 1, fetcher)
	}
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
	time.Sleep(time.Second)
}