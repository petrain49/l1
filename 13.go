package main

/*
Поменять местами два числа без создания временной переменной.
*/
func switchNums(arr []int, fi, si int) {
	arr[fi], arr[si] = arr[si], arr[fi]
}
