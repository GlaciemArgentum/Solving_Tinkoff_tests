package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// CheckWord проверяет, является ли данная последовательность цветов некрасивой.
func CheckWord(color string) int {
	pre := int32(color[0])
	for _, i := range color[1:] {
		if i == pre {
			return 1
		}
		pre = i
	}
	return 0
}

func main() {
	var (
		n int    // n — колличество букв.
		s string // s — искомая строка.
		b string // b — распределение цветов.
	)
	var countOfBadWords int
	colors := make([]string, 0, 2*n)

	fmt.Scan(&n)
	in := bufio.NewReader(os.Stdin)
	s, _ = in.ReadString('\n')
	b, _ = in.ReadString('\n')

	words := strings.Split(s, " ")
	counter := 0
	for _, i := range words {
		colors = append(colors, string([]rune(b)[counter:counter+len(i)]))
		counter += len(i)
	}
	for i := 0; i < len(words); i++ {
		countOfBadWords += CheckWord(colors[i])
	}
	fmt.Println(countOfBadWords)
}
