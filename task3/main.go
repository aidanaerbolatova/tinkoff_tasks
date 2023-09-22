package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var count int
	fmt.Fscan(in, &count)

	firstRow := make([]string, count)
	secondRow := make([]string, count)

	for i := 0; i < len(firstRow); i++ {
		fmt.Fscan(in, &firstRow[i])
	}

	for i := 0; i < len(secondRow); i++ {
		fmt.Fscan(in, &secondRow[i])
	}

	var index []int
	var temp1, temp2 []string

	for i := 0; i < count; i++ {
		if firstRow[i] != secondRow[i] {
			index = append(index, i)
		}
	}
	if len(index) >= 2 {
		for i := index[0]; i <= index[len(index)-1]; i++ {
			temp1 = append(temp1, firstRow[i])
			temp2 = append(temp2, secondRow[i])
		}
	}

	sort.Strings(temp1)

	if equal_slice(temp1, temp2) {
		fmt.Fprintln(out, "YES")
	} else {
		fmt.Fprintln(out, "NO")
	}
}

func equal_slice(arr1, arr2 []string) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for i, value := range arr1 {
		if value != arr2[i] {
			return false
		}
	}
	return true
}

// func main() {
// 	var count int
// 	in := bufio.NewReader(os.Stdin)
// 	fmt.Fscan(in, &count)

// 	scanner := bufio.NewScanner(os.Stdin)

// 	scanner.Scan()
// 	firstRow := scanner.Text()
// 	first := strings.Fields(firstRow)

// 	scanner.Scan()
// 	secondRow := scanner.Text()
// 	second := strings.Fields(secondRow)

// 	if len(first) != count || len(second) != count {
// 		fmt.Println("NO")
// 		return
// 	}

// 	var index []int
// 	var temp, temp2 []string
// 	for i := 0; i < len(first); i++ {
// 		if first[i] != second[i] {
// 			index = append(index, i)
// 		}
// 	}
// 	if len(index) < 2 {
// 		fmt.Println("YES")
// 		return
// 	}
// 	for i := index[0]; i <= index[len(index)-1]; i++ {
// 		temp = append(temp, first[i])
// 		temp2 = append(temp2, second[i])
// 	}

// 	sort.Strings(temp)

// 	if slice_equality(temp, temp2) {
// 		fmt.Println("YES")
// 	} else {
// 		fmt.Println("NO")
// 	}

// }

// func slice_equality(str1, str2 []string) bool {
// 	if len(str1) != len(str2) {
// 		return false
// 	}
// 	for i, str := range str1 {
// 		if str != str2[i] {
// 			return false
// 		}
// 	}
// 	return true
// }
