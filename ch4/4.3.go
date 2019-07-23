// 重写函数 reverse, 使用数组指针作为参数而不是 slice
package main

import "fmt"

func reverse(arr *[5]int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	reverse(&arr)
	fmt.Println(arr)
}
