// server2 是一个迷你的回声和计数器服务器
// 这个程序使用浏览器测试有一个问题，http://localhost:8000/ 会出发两次 count++ ，原因是默认会多请求一次 http://localhost:8000/favicon.ico
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler2)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler2(w http.ResponseWriter, r *http.Request)  {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Println("count++")
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request)  {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}


