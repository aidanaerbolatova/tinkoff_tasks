package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	scanner.Scan()
	row := scanner.Text()
	fmt.Println(row)
	rowArr := strings.Fields(row)

	spirit, _ := strconv.Atoi(rowArr[0])
	questions, _ := strconv.Atoi(rowArr[1])

	var temp []int
	var present [][]string
	var counter = make(map[string]int)

	for i := 1; i <= spirit; i++ {
		counter[strconv.Itoa(i)]++
		temp = append([]int{}, i)
		present = append(present, convertToString(temp))
	}

	for questions > 0 {
		questions--
		scanner.Scan()
		row := scanner.Text()
		rowElements := strings.Fields(row)

		switch rowElements[0] {
		case "1":
			present, counter = firstQuestion(present, rowElements, counter)
		case "2":
			fmt.Fprintln(out, secondQuestion(present, rowElements))
		case "3":
			fmt.Fprintln(out, thirdQuestion(rowElements, counter))
		}
	}
}

func convertToString(row []int) []string {
	var result []string
	for i := 0; i < len(row); i++ {
		result = append(result, strconv.Itoa(row[i]))
	}
	return result
}

func firstQuestion(present [][]string, row []string, counter map[string]int) ([][]string, map[string]int) {
	var result [][]string
	var new []string
	var flag int
	if len(row) == 3 {
		for _, value := range present {
			if !findItem(value, row[1]) && !findItem(value, row[2]) {
				result = append(result, value)
			} else {
				new = mergeArr(new, value)
				flag++
			}
		}
		if flag == 2 {
			result = append(result, new)
			for _, v := range new {
				counter[v]++
			}
		}
	}
	return result, counter
}

func secondQuestion(present [][]string, row []string) string {
	if len(row) == 3 {
		for _, value := range present {
			if findItem(value, row[1]) && findItem(value, row[2]) {
				return "YES"
			}
		}
	}
	return "NO"
}

func thirdQuestion(row []string, counter map[string]int) int {
	if len(row) == 2 {
		for k, v := range counter {
			if k == row[1] {
				return v
			}
		}
	}
	return 0
}

func findItem(row []string, item string) bool {
	for _, value := range row {
		if value == item {
			return true
		}
	}
	return false
}

func mergeArr(a, b []string) []string {
	var c []string
	c = append(a, b...)
	return c
}
