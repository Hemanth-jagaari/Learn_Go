package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("about server creation in golang")
	//PerformGetRequest()
	//PerformPostRequest()
	PerformPostFormRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"

	res, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	fmt.Println("Status code=", res.Status)
	fmt.Println("content length", res.ContentLength)

	content, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(content)
	fmt.Println(string(content))

	var resBuilder strings.Builder

	resBuilder.Write(content)

	fmt.Println(resBuilder.String())
}
func PerformPostRequest() {
	const myurl = "http://localhost:8000/post"

	requestBody := strings.NewReader(`
		{
			"name":"Hemanth",
			"age":9
		}
	
	`)
	res, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	content, _ := io.ReadAll(res.Body)
	var resBuilder strings.Builder

	fmt.Println("Status code=", res.Status)
	fmt.Println("content length ", res.ContentLength)
	resBuilder.Write(content)
	fmt.Println(resBuilder.String())
}

func PerformPostFormRequest() {
	const myurl = "http://localhost:8000/postform"

	data := url.Values{}

	data.Add("country", "india")
	data.Add("city", "hyderabad")
	data.Add("state", "telangana")
	res, err := http.PostForm(myurl, data)

	if err != nil {
		panic(err)
	}
	fmt.Println("Status code=", res.Status)
	defer res.Body.Close()
	content, _ := io.ReadAll(res.Body)
	fmt.Println(string(content))

}
