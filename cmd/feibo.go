package main

import
(
	"fmt"
	"time"
)

const LIM = 47

func main() {
	//result := 0
	//var array []int
	t := time.Now()
	fmt.Println(t)

	var array [LIM]int
	for i := 0; i < LIM; i++ {
		array[i] = fibonacci(i)
	}
	//fmt.Println(array)
	t2 := time.Now()
	fmt.Println(t2)
}

func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}