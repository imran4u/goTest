package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {

	links := []string{"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://youtube-wrong.com",
		"http://amazon.com",
		"http://wikipedia.org",
		"http://twitter.com",
	}

	startTime := time.Now()
	for _, link := range links {
		checkLink(link, nil)
	}
	endTime := time.Now()

	fmt.Println("Without goRoutine Total Time: ", endTime.Sub(startTime))

	// Using Go routine
	wg := &sync.WaitGroup{}
	wg.Add(len(links))
	startTime = time.Now()
	for _, link := range links {
		go checkLink(link, wg)
	}
	endTime = time.Now()

	wg.Wait()
	fmt.Println("With goRoutine Total Time: ", endTime.Sub(startTime))

}

func checkLink(link string, wg *sync.WaitGroup) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "is down!")
		if wg != nil {
			wg.Done()
		}
		return
	}
	fmt.Println(link, "is up!")
	if wg != nil {
		wg.Done()
	}
}
