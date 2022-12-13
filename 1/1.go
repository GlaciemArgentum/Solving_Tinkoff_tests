package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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
	var n, countOfBadWords int
	var s, b string
	//sR := make([]rune, 0, 32)
	colors := make([]string, 0, 2*n)

	fmt.Scan(&n)
	in := bufio.NewReader(os.Stdin)
	/*sR = fmt.Scan()
	scanner :=
	scanner.Scan()
	s = scanner.Text()*/
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

//Algorithms and Data Structures
