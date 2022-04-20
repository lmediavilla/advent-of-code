package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	//Open the file
	file, err := os.Open("input")
	if err != nil {
		os.Exit(0)
	}
	//Read the file
	fs := bufio.NewScanner(file)
	var numbersstring []string
	for fs.Scan() {
		numbersstring = append(numbersstring, fs.Text())
	}
	file.Close()
	var numbers = make([]int, len(numbersstring))
	res := 0
	for i, num := range numbersstring {
		numbers[i], err = strconv.Atoi(num)
		if err != nil {
			os.Exit(0)
		}
		if i > 0 && (numbers[i] > numbers[i-1]) {
			res++
		}
	}
	fmt.Println(res)
}
