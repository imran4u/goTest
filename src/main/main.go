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
	c := make(chan string)
	startTime := time.Now()

	for _, link := range links {
		go checkLink(link, c)
	}

	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}
	fmt.Println("All chanels closed")

	endTime := time.Now()

	fmt.Println("Total Time: ", endTime.Sub(startTime))
	// time.Sleep(20 * time.Second)

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		c <- fmt.Sprintf("%s: %s", link, "is Down!!")

		return
	}
	c <- fmt.Sprintf("%s, %s", link, "is up!!")
}
