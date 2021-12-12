package main

import (
	"strings"
)

/*
Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».
*/
func rvsString(s string) string {

	sas := strings.Split(s, " ")
	rvsWords(sas)
	res := strings.Join(sas, " ")

	return res
}

func rvsWords(sas []string) {
	for l, r := 0, len(sas)-1; l < r; l, r = l+1, r-1 {
		sas[l], sas[r] = sas[r], sas[l]
	}
}
