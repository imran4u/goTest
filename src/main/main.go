package main

import (
	"fmt"
	"net/http"
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
		go checkLink(link)
	}

	endTime := time.Now()

	fmt.Println("Total Time: ", endTime.Sub(startTime))
	time.Sleep(20 * time.Second)

}

func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "is down!")
		return
	}
	fmt.Println(link, "is up!")
}
