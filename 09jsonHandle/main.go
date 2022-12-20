package main

import (
	"encoding/json"
	"fmt"
)

type Course struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"price"`
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("json data handling")
	//EncodeJson()
	DecodeJson()
}
func EncodeJson() {
	courses := []Course{
		{"React", 1234, "LCO", "123ABC", []string{"JS", "Node"}},
		{"Golang", 2345, "Lco", "234ABC", []string{"GO", "AWS"}},
		{"Angular", 1235, "LCO", "523ABC", nil},
	}

	fmt.Println(courses)
	finalJson, err := json.MarshalIndent(courses, "", "\t")

	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "Golang",
		"price": 2345,
		"website": "Lco",
		"tags": [
				"GO",
				"AWS"
		]
	}
	`)

	check := json.Valid(jsonDataFromWeb)

	if check {
		fmt.Println("valid json data")
	} else {
		fmt.Println("Not a valid json data")
	}
	var course Course
	json.Unmarshal(jsonDataFromWeb, &course)

	fmt.Printf("%#v\n", course)
	fmt.Println(course)

	var onlineData map[string]interface{}

	json.Unmarshal(jsonDataFromWeb, &onlineData)
	fmt.Printf("%#v\n", onlineData)

	fmt.Println(onlineData)

	for k, v := range onlineData {
		fmt.Printf("key=%v and value=%v and Type:%T\n", k, v, v)
	}
}
