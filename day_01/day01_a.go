package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
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

	// Sort lists
	sort.Slice(a_list, func(i, j int) bool {
		return a_list[i] < a_list[j]
	})
	sort.Slice(b_list, func(i, j int) bool {
		return b_list[i] < b_list[j]
	})

	fmt.Println(a_list)
	fmt.Println(b_list)

	var result int64 = 0
	for i := range a_list {
		tmp := int64(math.Abs(float64(a_list[i]) - float64(b_list[i])))
		fmt.Printf("Value of math.Abs: %d \n", tmp)
		result = result + tmp
	}
	print(result)

}
