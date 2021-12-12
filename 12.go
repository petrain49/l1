package main

/*
Имеется последовательность строк - (cat, cat, dog, cat, tree).
Создать для нее собственное множество.
*/

type stringSet map[string]int

func (s *stringSet) add(str string) {
	_, ok := (*s)[str]
	if !ok {
		(*s)[str] = 1
	} else {
		(*s)[str]++
	}
}
