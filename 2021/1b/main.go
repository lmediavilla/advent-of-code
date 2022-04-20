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
	res1, res2, total := 0, 0, 0
	for i, num := range numbersstring {
		numbers[i], err = strconv.Atoi(num)
		if err != nil {
			os.Exit(0)
		}
	}
	for i, num := range numbers {
		if i+3 < len(numbers) {
			res1 = num + numbers[i+1] + numbers[i+2]
			res2 = res1 - num + numbers[i+3]
			if res2 > res1 {
				total++
			}
		}
	}
	fmt.Println(total)
}
