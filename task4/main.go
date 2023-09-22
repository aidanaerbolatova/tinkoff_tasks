package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	cash, count := 0, 0
	fmt.Fscanf(in, "%d %d", &cash, &count)

	array := make([]int, count*2)

	for i := 0; i < len(array); i += 2 {
		var temp int
		fmt.Fscan(in, &temp)
		array[i] = temp
		array[i+1] = temp
	}

	result := backtracking([]int{}, array, cash, 0)
	if result == nil {
		fmt.Fprintln(out, -1)
	} else {
		fmt.Fprintln(out, len(result))
		for index, element := range result {
			fmt.Fprint(out, element)
			if index < len(result)-1 {
				fmt.Fprint(out, " ")
			}
		}
		fmt.Fprintln(out)
	}

}

func backtracking(new []int, array []int, cash int, start int) []int {
	if cash < 0 {
		return nil
	} else if cash == 0 {
		return new
	}

	var result []int
	for i := start; i < len(array); i++ {
		if i > start && array[i] == array[i-1] {
			continue
		}
		new = append(new, array[i])
		template := backtracking(new, array, cash-array[i], i+1)
		if template != nil {
			if result == nil || len(template) > len(result) {
				result = append([]int{}, template...)
			}
		}
		new = new[:len(new)-1]
	}

	return result
}
