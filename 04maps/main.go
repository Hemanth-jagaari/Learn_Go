package main

import "fmt"

type Person struct {
	name   string
	age    int
	salary float32
}

var val int = 123

func change(person *Person, m map[string]string) {
	person.name = "vinay"
	val = 128
	fmt.Println("inner value", val)
	fmt.Println("in change func")
	for val, key := range m {
		fmt.Printf("key %v -> value %v\n", val, key)
	}
	m["hello"] = "hi"
}
func fib(i int) int {
	if i <= 1 {
		return i
	} else {
		return fib(i-1) + fib(i-2)
	}
}
func check(arr []int) {
	for val := range arr {
		fmt.Println(arr[val])
		arr[val] = -1
	}
}
func main() {
	fmt.Println("hello in main")
	p := Person{age: 12, name: "hemanth", salary: 123.6}
	fmt.Println(p)
	m := map[string]string{"class": "A"}
	m["name"] = "hemanth"
	m["first"] = "jagaari"
	m["last"] = "hemanth"
	m["branch"] = "cse"
	fmt.Println(m)
	change(&p, m)
	fmt.Println(p)
	fmt.Println(m)
	fmt.Println("fib 5", fib(4))
	brr := make([]int, 4)
	brr = append(brr, 12, 34, 12, 3, 4, 4, 4)

	check(brr)

	//check(brr)
	c := [6]int{1, 23213}
	fmt.Println(c)
	arr := make([][6]int, 3)
	arr[0][0] = 1
	arr[1][1] = 2
	arr[2][2] = 3
	fmt.Println(arr)

}
