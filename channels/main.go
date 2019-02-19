package main

import (
	"fmt"
	"net/http"
)

func main() {

	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://bomnegocio.com",
		"http://golang.org",
		"http://amazon.com",
	}

	for _, link := range links {
		go checkLink(link)
	}
}

func checkLink(link string) {
	_, error := http.Get(link)
	if error != nil {
		fmt.Println(link, " Must be down!")
		return
	}

	fmt.Println(link, " is up an running!")
}
