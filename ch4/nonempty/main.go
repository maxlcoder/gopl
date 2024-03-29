// Nonempty 演示了 slice 的就地修改算法
package main

import "fmt"

// nonempty 返回一个新的 slice, slice 中的元素都是非空字符串
// 在函数的调用过程中，底层数组的元素发生了改变
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func nonempty2(strings []string) []string {
	out := strings[:0] // 引用原始的 slice 的新的零长度的 slice
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

// 第二种顺序不为持的做法
// func remove(slice []int, i int) []int {
// 	slice[i] = slice[len(slice)-1]
// 	return slice[:len(slice)-1]
// }

func main() {
	s := []int{5, 6, 7, 8, 9}
	fmt.Println(remove(s, 2)) // "[5 6 8 9]"
}
