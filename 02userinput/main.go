package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("welcome here")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter Rating of Pizza :")
	input, _ := reader.ReadString('\n')
	fmt.Println("Rating of Pizza :", input)
	rating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(rating + 1)
	}

}
