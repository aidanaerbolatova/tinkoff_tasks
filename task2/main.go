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

	var text string
	fmt.Fscan(in, &text)

	var all = make(map[rune]int)
	for _, s := range text {
		all[s]++
	}
	count := 0
	flag := true
	for flag {
		for _, v := range "sheriff" {
			if all[v] == 0 {
				flag = false
				break
			}
			all[v]--
		}
		if flag {
			count++
		}
	}
	fmt.Fprintln(out, count)
}
