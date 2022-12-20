package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://lco.dev:3000/learn?coursename=golang&intructor=hitesh"

func main() {
	fmt.Println("about url handling")

	fmt.Println("url is", myurl)
	res, err := url.Parse(myurl)
	if err != nil {
		panic(err)
	}
	fmt.Println(res.Scheme)
	fmt.Println(res.Host)
	fmt.Println(res.Port())
	fmt.Println(res.RawQuery)
	fmt.Println(res.Query())
	partsorurl := &url.URL{
		Scheme:   "https",
		Host:     "lco.dev:3000",
		Path:     "/learn",
		RawQuery: "course=golang",
	}
	fmt.Println(partsorurl.String())
}
