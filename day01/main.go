package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)
	var col1 []int
	var col2 []int

	for scanner.Scan() {
		line := scanner.Text()
		split_array := strings.Split(line, " ")
		a, _ := strconv.Atoi(split_array[0])
		b, _ := strconv.Atoi(split_array[1])

		col1 = append(col1, a)
		col2 = append(col2, b)
	}

	sort.Slice(col1, func(i, j int) bool {
		return col1[i] < col1[j]
	})

	sort.Slice(col2, func(i, j int) bool {
		return col2[i] < col2[j]
	})

	sum := 0

	for i := 0; i < len(col1); i++ {
		x1, x2 := col1[i], col2[i]
		difference := 0
		if x1 < x2 {
			difference = x2 - x1
		} else {
			difference = x1 - x2
		}

		sum += difference
	}

	score := 0

	for i := 0; i < 1000; i++ {
		similar := 0
		for j := 0; j < 1000; j++ {
			if col1[i] == col2[j] {
				similar++
			}
		}

		score = score + (col1[i] * similar)
	}

	println(score)
}
