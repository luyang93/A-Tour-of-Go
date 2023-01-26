package main

import (
	"fmt"
	"sync"
)

type SafeUrlMap struct {
	mu sync.Mutex
	v  map[string]bool
}

func (s *SafeUrlMap) checkVisited(url string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, ok := s.v[url]
	if ok == false {
		s.v[url] = true
		return false
	}
	return true
}

var s = SafeUrlMap{v: make(map[string]bool)}

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, ch chan bool) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	if depth <= 0 || s.checkVisited(url) {
		ch <- false
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		ch <- false
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	chs := make(map[string]chan bool, len(urls))
	for _, u := range urls {
		chs[u] = make(chan bool)
		go Crawl(u, depth-1, fetcher, chs[u])
	}
	for _, v := range urls {
		<-chs[v]
	}
	ch <- true
	return
}

func main() {
	ch := make(chan bool)
	go Crawl("https://golang.org/", 4, fetcher, ch)
	<-ch
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
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
