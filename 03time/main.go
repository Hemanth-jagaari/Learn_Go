package main

import (
	"fmt"
	"sort"
	"time"
)

func main() {
	fmt.Println("Welcome")
	timenow := time.Now()
	fmt.Println("Currnt time: ", timenow)
	fmt.Println(timenow.Format("01-02-2006 15:04:05 Monday"))

	var arr = []int{1, 2, 3, 4}
	fmt.Println(arr)
	var brr = []int{}
	fmt.Println(brr)
	brr = append(brr, 4, 2, 1, 5)
	fmt.Println(brr)
	fmt.Println(sort.IntsAreSorted(brr))
	sort.Ints(brr)
	fmt.Println(brr)
}
