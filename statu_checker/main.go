package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
	}
	c := make(chan string)

	for _, link := range links {
		go checkline(link, c)
	}

	// main go routine will exit if we dont add a check to wait for the othe go routines to finish the execution
	// for {
	// 	go checkline(<-c, c)
	// }

	for l := range c { //same thing as above for loop
		time.Sleep(5 * time.Second)
		go func() {
			time.Sleep(5 * time.Second)
			checkline(l, c)
		}()
	}
}

func checkline(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "may be down")
		c <- link
		// return
	}
	fmt.Println(link, "is up")
	c <- link
}
