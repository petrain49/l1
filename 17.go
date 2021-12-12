package main

/*
Реализовать бинарный поиск встроенными методами языка.
*/
func binSearch(elem int, m []int) bool {
	l := 0
	h := len(m) - 1
	mdn := 0

	for l <= h {
		mdn = (l + h) / 2
		if m[mdn] < elem {
			l = mdn + 1
			continue
		}
		h = mdn - 1

	}

	if l == len(m) || m[l] != elem {
		return false
	}

	return true
}
