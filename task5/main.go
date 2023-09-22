package main

import (
	"bufio"
	"fmt"
	"os"
)

type pathway struct {
	city   int
	length int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var countCity, countRoad int
	fmt.Fscan(in, &countCity, &countRoad)

	all := make([][]pathway, countCity+1)

	for i := 0; i < countRoad; i++ {
		var v, u, w int
		fmt.Fscan(in, &v, &u, &w)

		all[v] = append(all[v], pathway{city: u, length: w})
		all[u] = append(all[u], pathway{city: v, length: w})
	}

	from := 1
	to := int(1e9) + 1
	var result, min int

	states := connect(all, countCity, min)

	for from <= to {
		midway := (from + to) / 2
		current := connect(all, countCity, midway)
		if current == states {
			result = max(result, midway)
			from = midway + 1
		} else {
			from = midway - 1
		}
	}

	fmt.Fprintln(out, result)
}

func connect(all [][]pathway, countCity, min int) int {
	result := 0
	visit := make([]bool, countCity+1)

	for i := 1; i <= countCity; i++ {
		if !visit[i] {
			result++
			checkRoad(all, visit, i, min)
		}
	}

	return result
}

func checkRoad(all [][]pathway, visit []bool, fromCity, min int) {
	for _, road := range all[fromCity] {
		if !visit[road.city] && road.length > min {
			visit[road.city] = true
			checkRoad(all, visit, road.city, min)
		}
	}
}

func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}

// type pathway struct {
// 	city   int
// 	length int
// }

// func main() {
// 	scanner := bufio.NewScanner(os.Stdin)
// 	scanner.Split(bufio.ScanWords)

// 	scanner.Scan()
// 	countCity, _ := strconv.Atoi(scanner.Text())

// 	scanner.Scan()
// 	countRoad, _ := strconv.Atoi(scanner.Text())

// 	all := make([][]pathway, countCity+1)

// 	for i := 0; i < countRoad; i++ {
// 		scanner.Scan()
// 		v, _ := strconv.Atoi(scanner.Text())

// 		scanner.Scan()
// 		u, _ := strconv.Atoi(scanner.Text())

// 		scanner.Scan()
// 		w, _ := strconv.Atoi(scanner.Text())

// 		all[v] = append(all[v], pathway{city: u, length: w})
// 		all[u] = append(all[u], pathway{city: v, length: w})
// 	}

// 	from := 1
// 	to := int(1e9) + 1
// 	result := 0
// 	min := 0
// 	states := connect(countCity, all, min)

// 	for from <= to {
// 		midway := (from + to) / 2
// 		current := connect(countCity, all, midway)
// 		if current == states {
// 			result = max(result, midway)
// 			from = midway + 1
// 		} else {
// 			to = midway - 1
// 		}
// 	}

// 	fmt.Println(result)
// }

// func connect(countCity int, all [][]pathway, min int) int {
// 	result := 0
// 	check := make([]bool, countCity+1)

// 	for i := 1; i <= countCity; i++ {
// 		if !check[i] {
// 			result++
// 			checkRoad(i, all, check, min)
// 		}
// 	}

// 	return result
// }

// func checkRoad(at int, all [][]pathway, check []bool, min int) {
// 	for _, d := range all[at] {
// 		if !check[d.city] && d.length > min {
// 			check[d.city] = true
// 			checkRoad(d.city, all, check, min)
// 		}
// 	}
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }
