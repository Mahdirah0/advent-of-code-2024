package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// condition to be safe
// all must be increasing or decreasing
// increasing or decreasing by 3

func get_difference(num1 int, num2 int) int {
	if num1 < num2 {
		return num2 - num1
	}

	return num1 - num2
}

func is_line_is_safe(is_increasing bool, val1 int, val2 int) bool {
	if val1 == val2 {
		return false
	}

	if get_difference(val1, val2) > 3 {
		return false
	}

	if is_increasing && val1 > val2 {
		return false
	}

	if !is_increasing && val1 < val2 {
		return false
	}

	return true
}

func copy_array(base_array []int, skip_index int) []int {
	copy := []int{}
	for i := 0; i < len(base_array); i++ {
		if skip_index == i {
			continue
		}
		copy = append(copy, base_array[i])
	}
	return copy
}

func check_line(line []int, is_inc bool, m int) bool {
	left := 0
	right := 1
	size := len(line)

	if m == 0 {
		return false
	}

	for right < size {
		val1 := line[left]
		val2 := line[right]

		if !is_line_is_safe(is_inc, val1, val2) {
			copy1 := copy_array(line, left)
			copy2 := copy_array(line, right)
			ma := m - 1

			if check_line(copy1, is_inc, ma) {
				return true
			} else if check_line(copy2, is_inc, ma) {
				return true
			}
			return false
		}

		left++
		right++
	}

	return true
}

func main() {
	file, err := os.Open("sample")
	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)
	safe_reports := 0
	lines_string := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		lines_string = append(lines_string, line)
	}

	for i := 0; i < len(lines_string); i++ {
		line_split := strings.Split(lines_string[i], " ")
		line := []int{}

		for k := 0; k < len(line_split); k++ {
			num, _ := strconv.Atoi(line_split[k])
			line = append(line, num)
		}

		first_val := line[0]
		second_val := line[1]
		is_inc := first_val < second_val

		if check_line(line, is_inc, 2) {
			safe_reports++
		}
	}

	fmt.Println(safe_reports)
}
