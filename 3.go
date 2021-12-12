package main

/*
Дана последовательность чисел: 2,4,6,8,10.
Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.
*/
func sum2(nums []int) int {

	num := make(chan int)
	defer close(num)

	res := 0
	for _, n := range nums {
		go pow2(n, num)
		res += <-num
	}

	return res
}
