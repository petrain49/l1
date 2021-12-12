package main

/*
Реализовать пересечение двух неупорядоченных множеств.
*/

type set map[int]bool

func (s *set) contains(n int) bool {
	_, ok := (*s)[n]

	return ok
}

func (s *set) intersection(otherSet set) set {
	resultSet := make(set)

	for n := range *s {
		if otherSet.contains(n) {
			resultSet[n] = true
		}
	}

	return resultSet
}
