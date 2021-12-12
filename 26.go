package main

import (
	"strings"
)

/*
Разработать программу, которая проверяет,
что все символы в строке уникальные (true — если уникальные, false etc).
Функция проверки должна быть регистронезависимой.
Например:
abcd — true
abCdefAaf — false
    aabcd — false

*/
func check(str string) bool {
	if len(str) < 2 {
		return true
	}
	str = strings.ToLower(str)

	runes := []rune(str)
	set := make(map[rune]bool)

	for _, r := range runes {
		_, ok := set[r]
		if !ok {
			set[r] = true
		} else {
			return false
		}
	}
	return true
}