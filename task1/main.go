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

	count, money := 0, 0
	fmt.Fscanf(in, "%d %d", &count, &money)
	var temp int
	for i := 0; i < count; i++ {
		var n int
		fmt.Fscan(in, &n)
		if n > temp && n < money {
			temp = n
		}
	}
	fmt.Fprintln(out, temp)

}

// func main() {
// 	scanner := bufio.NewScanner(os.Stdin)
// 	scanner.Scan()
// 	input := scanner.Text()
// 	parts := strings.Fields(input)
// 	count, _ := strconv.Atoi(parts[0])
// 	s, _ := strconv.Atoi(parts[1])
// 	scanner.Scan()
// 	input = scanner.Text()
// 	revolverPrices := strings.Fields(input)
// 	var prices, result []int
// 	for _, priceStr := range revolverPrices {
// 		if len(prices) == count {
// 			break
// 		}
// 		price, _ := strconv.Atoi(priceStr)
// 		prices = append(prices, price)
// 	}
// 	sort.Ints(prices)
// 	for _, nums := range prices {
// 		if nums < s {
// 			result = append(result, nums)
// 		}
// 	}
// 	if len(result) > 0 {
// 		fmt.Println(result[len(result)-1])
// 	} else {
// 		fmt.Println("0")
// 	}
// }
