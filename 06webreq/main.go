package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://lco.dev/"

func main() {
	fmt.Println("about web request")
	res, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	fmt.Println("response is", res.Body)
	defer res.Body.Close()
	databyte, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}
	fmt.Println("response contents", string(databyte))

}
