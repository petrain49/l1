package main

/*
Разработать программу, которая переворачивает подаваемую на ход строку,
например: «главрыба — абырвалг».
Символы могут быть unicode.
*/
func rvs(s string) string {
	if len(s) < 2 {
		return s
	}

	rs := []rune(s)

	for l, r := 0, len(rs)-1; l < r; l, r = l+1, r-1 {
		rs[l], rs[r] = rs[r], rs[l]
	}

	return string(rs)
}
