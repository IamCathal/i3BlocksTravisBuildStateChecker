package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func getBuildState(url string) {
	if len(url) < 1 {
	fmt.Println(" No url given")
	return
	}
	res, err := http.Head(url)
	if err != nil {
		fmt.Println(" Error getting page")
	}
	res, err = http.Get(url)
	if err != nil {
		fmt.Println(" Error getting page")
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(" Error reading response")
	}

	matched, err := regexp.Match(`passing`, body)
	if matched {
		fmt.Printf(" Passing")
	} else {
		fmt.Printf(" Failing")
	}
}

func main() {
	url := ""
	getBuildState(url)
}
