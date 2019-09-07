// 实现一次遍历即可完成元素旋转
package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	rotate(arr, 3)
	fmt.Println(arr)
}

func rotate(arr []int, n int) {
	if n > 5 {
		return
	}
	// 单独提取旋转的前几位
	var temp []int
	for _, v := range arr[:n] {
		temp = append(temp, v)
	}
	copy(arr, arr[n:])
	copy(arr[len(arr)-n:], temp)
}
