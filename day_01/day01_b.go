package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	fscanner := bufio.NewScanner(file)
	var a_list []int64
	var b_list []int64
	for fscanner.Scan() {
		pairs := strings.Split(fscanner.Text(), "   ")

		a, _ := strconv.ParseInt(pairs[0], 10, 64)
		b, _ := strconv.ParseInt(pairs[1], 10, 64)
		a_list = append(a_list, a)
		b_list = append(b_list, b)
	}

	lookup := make(map[int64]int64)
	for _, value := range b_list {

		if _, exists := lookup[value]; exists {
			lookup[value] += 1
		} else {
			lookup[value] = 1
		}
	}

	var result int64 = 0
	for _, value := range a_list {
		mult_val := lookup[value]
		result = result + value*mult_val
	}

	fmt.Println(result)
}
