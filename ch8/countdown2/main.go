package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	abort := make(chan struct{})

	// 此处进入并行，这个匿名函数中等待输入，一旦有输入了，立即发送终止信号到 abort channel
	go func() {
		os.Stdin.Read(make([]byte, 1)) // 读取单个字节
		abort <- struct{}{}
	}()

	// 由于上面的进入并行状态，这里的顺序执行了，后面的 select 等待 定时channel 和 abort channel 的输入返回也就是 轮询等待
	fmt.Println("Commencing countdown. Press return to abort.")
	select {
	case <-time.After(3 * time.Second):
		// 不执行任何操作
		fmt.Println("time less")
	case <-abort:
		fmt.Println("Launch aborted!")
		return
	}
	launch()
}

func launch() {
	fmt.Println("Lift off!")
}