package main

import (
	"encoding/hex"
	"fmt"
)

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
		a = append(a[:index], append([]int{num}, a[index:]...)...)
		return a
	}
	result := make([]int, len(a)+1, cap(a)*2)
	copy(result, a[:index])
	copy(result[index:], []int{num})
	copy(result[index+1:], a[index:])
	return result
}

func myDelete(a []int, index int) []int {
	if index > len(a) || index < 0 {
		return a
	}
	return append(a[:index], a[index+1:]...)
}

func main() {
	// fibonacci 验证，用的递归
	fmt.Println(fibonacci(5))

	// 验证 add and myDelete
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(a)
	fmt.Println(len(a), cap(a))

	b := add(a, 11, 4)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(len(a), cap(a))

	c := myDelete(a, 4)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(len(a), cap(a))

	// 验证序列化输出
	f := 698.123456
	fmt.Println(f)
	fmt.Printf("%.2f\n", f)

	d := []byte(`go进阶训练营`)
	fmt.Println(d)
	d1 := hex.EncodeToString(d)
	fmt.Println(d1)
	fmt.Println(hex.DecodeString(d1))
	fmt.Printf("%v\n", d)
	fmt.Printf("%#v\n", d)
	fmt.Printf("%s\n", d)
}
