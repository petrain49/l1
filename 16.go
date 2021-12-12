package main

/*
Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
*/
func quicksort(m []int) []int {
	if len(m) < 2 {
		return m
	}

	l, r := 0, len(m)-1
	pivot := len(m) / 2

	m[pivot], m[r] = m[r], m[pivot]

	for x := range m {
		if m[x] < m[r] {
			m[l], m[x] = m[x], m[l]
			l++
		}
	}

	m[l], m[r] = m[r], m[l]

	quicksort(m[:l])
	quicksort(m[l+1:])

	return m
}
