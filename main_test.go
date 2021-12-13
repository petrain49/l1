package main

import (
	"sync"
	"testing"
)

func TestOne(t *testing.T) {
	one()
}

func TestTwo(t *testing.T) {
	two()
}

func TestThree(t *testing.T) {
	nums := []int{2, 4, 6, 8, 10}
	res := sum2(nums)

	t.Log(res)
	if res != 220 {
		t.FailNow()
	}
}

/*
func TestFour(t *testing.T) {
	work(45)
}
*/

func TestFive(t *testing.T) {
	wnr(2)
}

func TestSix(t *testing.T) {
	six()
}

func TestSeven(t *testing.T) {
	res := seven()
	t.Log(res)
}

func TestEight(t *testing.T) {
	// 45 = 101101
	var num int64 = 45

	// 61 = 111101
	res1 := setBit(num, 4, true)
	t.Log(res1)
	if res1 != 61 {
		t.Fail()
	}

	// 13 = 1101
	res0 := setBit(num, 5, false)
	t.Log(res0)
	if res0 != 13 {
		t.Fail()
	}
}

func TestNine(t *testing.T) {
	pipeline()
}

func TestTen(t *testing.T) {
	arr := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	t.Log(combineTemps(arr))
}

func TestEleven(t *testing.T) {
	first := set{ 1:true, 2:true, 3:true, 4:true, 5:true }
	second := set{ 3:true, 4:true, 5:true, 6:true, 7:true }

	t.Log(first.intersection(second))
}

func TestTwelve(t *testing.T) {
	arr := [...]string{ "cat", "cat", "dog", "cat", "tree" }
	set := make(stringSet)

	for _, s := range arr {
		set.add(s)
	}

	t.Log(set)
}

func TestThirteen(t *testing.T) {
	arr := []int{0, 1, 2, 3}

	switchNums(arr, 1, 2)

	t.Log(arr)
}

func TestFourteen(t *testing.T) {
	ch := make(chan int)
	if typeDet(ch) != "chan" {
		t.FailNow()
	}

	n := 0
	if typeDet(n) != "int" {
		t.FailNow()
	}

	s := ""
	if typeDet(s) != "string" {
		t.FailNow()
	}

	b := false
	if typeDet(b) != "bool" {
		t.FailNow()
	}
}

func TestFifteen(t *testing.T) {
	someFunc()
}

func TestSixteen(t *testing.T) {
	m := []int{ 10, 9, 8, 7, 6, 50, 4, 3, 2, 1, 0 }
	t.Log(quicksort(m))
}

func TestSeventeen(t *testing.T) {
	m := []int{ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10 }

	if binSearch(6, m) != true {
		t.FailNow()
	}
	if binSearch(-2, m) != false {
		t.FailNow()
	}
}

func TestEighteen(t *testing.T) {
	var wg sync.WaitGroup

	counter := new(counter)

	for x := 0; x < 10000; x++ {
		wg.Add(1)
		go func() {
			counter.inc()
			wg.Done()
		}()
	}
	wg.Wait()
	
	t.Log(counter.num)

	if counter.num != 10000 {
		t.FailNow()
	}
}

func TestNineteen(t *testing.T) {
	s := "главрыба🐸"
	res := rvs(s)
	t.Log(res)

	if res != "🐸абырвалг" {
		t.FailNow()
	}
}

func TestTwenty(t *testing.T) {
	s := "snow dog sun"

	res := rvsString(s)
	if res != "sun dog snow" {
		t.FailNow()
	}
}

func TestTwentyOne(t *testing.T) {
	s := "hello"

	adapter := adapter{service: &XMLData{}, data: s}
	res := adapter.handleJSONData()

	t.Log(res)

	if res != "Converted to json: Adapter: hello" {
		t.FailNow()
	}
}

func TestTwentyTwo(t *testing.T) {
	res := input("24", "+", "25")
	t.Log(res)

	if res != "49" {
		t.FailNow()
	}
}

func TestTwentyThree(t *testing.T) {
	s := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	n := 5
	delete(s, n)

	t.Log(s)
}

func TestTwentyFour(t *testing.T) {
	f := newPoint(4, 3)
	s := newPoint(0, 0)
	res := distance(f, s)

	t.Log(res)
	if res != 5 {
		t.FailNow()
	}
}

func TestTwentyFive(t *testing.T) {
	t.Log(mySleepAfter(5))
	t.Log(mySleepSubstraction(5))
	t.Log(mySleepTimer(5))
}

func TestTwentySix(t *testing.T) {
	t.Log(check("abcd"),
		check("abCdefA"),
		check("aabcd"),
		check("🐸🐸"),
	)
}
