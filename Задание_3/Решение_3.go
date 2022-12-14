package main

import "fmt"

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	var n int
	_, err := fmt.Scan(&n)
	CheckErr(err)

	for i := 2; ; i++ {
		if n%i == 0 {
			fmt.Println(n/i, n-n/i)
			break
		}
	}
}
