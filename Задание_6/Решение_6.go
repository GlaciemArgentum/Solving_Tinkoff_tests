package main

import "fmt"

// maxXOR возвращает максимальное значение исключающего ИЛИ введённого числа k и каждого элемента list.
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
	var (
		n      int // количество запросов.
		maxPre int // запоминаем предыдущее максимальное значение (см. описание ф-ции maxXOR).
	)
	list := make(map[int]int)
	fmt.Scan(&n)

	var k int // добавочное число.
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
