package main

import (
	"fmt"
)

type Currency struct {
	sum, values int
}

// CheckErr проверяет входные данные на ошибку.
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

// Contains проверяет наличине тройки значений в проверочном массиве.
func Contains(a [][3]int, x [3]int) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func main() {
	var (
		err    error
		wallet [3]Currency
		base   [3]int
	)
	excludeList := make([][3]int, 0, 32) // проверочный массив уже учтённых троек значений.
	var absMoney int                     // сумма всех денег в общей системе единиц.

	// считываем курсы валют.
	for i := 0; i < 3; i++ {
		_, err = fmt.Scan(&wallet[i].values)
		CheckErr(err)
	}

	// считываем колличество денег разных валют.
	for i := 0; i < 3; i++ {
		_, err = fmt.Scan(&wallet[i].sum)
		CheckErr(err)
	}

	// считаем absMoney
	for i := 0; i < 3; i++ {
		base[i] = wallet[i].sum % wallet[i].values
		absMoney += (wallet[i].sum - base[i]) / wallet[i].values
	}

	variable := [3]int{0, 0, 0}

	// перебор всех возможных вариантов троек значение.
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
