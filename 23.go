package main

/*
Удалить i-ый элемент из слайса.
*/
func delete(s []string, i int) {
	s[i], s[len(s)-1] = s[len(s)-1], ""
	s = s[:len(s)-1]
}
