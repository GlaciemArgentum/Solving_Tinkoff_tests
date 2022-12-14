package main

import (
	"fmt"
)

type Currency struct {
	sum, values int
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func Contains(a [][3]int, x [3]int) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func main() {
	var err error
	var wallet [3]Currency
	var base [3]int
	excludeList := make([][3]int, 0, 32)
	var absMoney int
	for i := 0; i < 3; i++ {
		_, err = fmt.Scan(&wallet[i].values)
		CheckErr(err)
	}
	for i := 0; i < 3; i++ {
		_, err = fmt.Scan(&wallet[i].sum)
		CheckErr(err)
	}

	for i := 0; i < 3; i++ {
		base[i] = wallet[i].sum % wallet[i].values
		absMoney += (wallet[i].sum - base[i]) / wallet[i].values
	}

	variable := [3]int{0, 0, 0}
	for i := 0; i <= absMoney; i++ {

		for j := 0; j <= absMoney-i; j++ {
			variable = [3]int{i + base[0], j + base[1], absMoney - i - j + base[2]}
			if !Contains(excludeList, variable) {
				excludeList = append(excludeList, variable)
			}
		}
	}
	fmt.Println(len(excludeList))
}
