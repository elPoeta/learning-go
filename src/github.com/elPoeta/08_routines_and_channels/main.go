package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// link, err := http.Get("https://google.com")
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// fmt.Println(link)
	sites := []string{
		"https://github.com",
		"https://google.com",
		"https://stackoverflow.com",
		"https://golang.org",
	}
	fmt.Println(sites)
	c := make(chan string)
	for _, link := range sites {
		go checkStatus(link, c)
	}
	//fmt.Println(<-c)
	for l := range c {
		go func(link string) {
			time.Sleep(time.Hour)
			checkStatus(link, c)
		}(l)
	}
}

func checkStatus(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println("[" + link + " ] STATUS : DOWN!")
		c <- link
	}
	fmt.Println("[ " + link + " ] STATUS : UP!")
	c <- link
}
