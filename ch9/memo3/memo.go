// memo包提供了一个对类型 Func 并发不安全的函数记忆功能
package memo

import "sync"

// Memo 缓存了调用 Func 的结果
type Memo struct {
	f Func
	mu sync.Mutex // 保护 cache
	cache map[string]result
}

// Func 是用于记忆的函数类型
type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err error
}

func New(f Func) *Memo  {
	return &Memo{f: f, cache: make(map[string]result)}
}

// Get 是并发安全的，阻塞缓存
func (memo *Memo) Get(key string) (interface{}, error)  {
	memo.mu.Lock() // 第一次上锁
	res, ok := memo.cache[key]
	memo.mu.Unlock() // 第一次解锁
	if !ok {
		res.value, res.err = memo.f(key)
		// 在两个临界区域之前，可能会有多个 goroutine 来计算 f(key) 并且
		// 更新 map
		memo.mu.Lock() // 第二次上锁
		memo.cache[key] = res
		memo.mu.Unlock() // 第二次解锁
	}
	return res.value, res.err
}


