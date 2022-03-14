package main

import "fmt"

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	return fibonacci(n-1) + fibonacci(n-2)
}

func add(a []int, num, index int) []int {
	if index > len(a) || index < 0 {
		return a
	}
	// 如果不用触发扩容，则可以在原切片上返回，避免资源的浪费
	if len(a)+1 < cap(a) {
		return append(a[:index], append([]int{num}, a[index:]...)...)
	}
	result := make([]int, 0, cap(a)*2)
	copy(result, a[:index])
	copy(result, []int{num})
	copy(result, a[index:])
	return result
}

func delete(a []int, index int) []int {
	if index > len(a) || index < 0 {
		return a
	}
	return append(a[:index], a[index+1:]...)
}

func main() {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(a)
	fmt.Println()
	fmt.Println(fibonacci(20))
}
