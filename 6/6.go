package main

import "fmt"

func maxXOR(list map[int]int, k int) int {
	if len(list) == 0 {
		return 0
	}
	max := 0
	xor := 0
	for key := range list {
		xor = k ^ key
		if xor > max {
			max = xor
		}
	}
	return max
}

func main() {
	var n, maxPre int
	list := make(map[int]int)
	fmt.Scan(&n)

	var k int
	for i := 0; i < n; i++ {
		fmt.Scan(&k)

		if _, inMap := list[k]; inMap {
			fmt.Println(maxPre)
		} else {
			if max := maxXOR(list, k); max > maxPre {
				maxPre = max
			}
			fmt.Println(maxPre)
			list[k] = k
		}
	}
}
