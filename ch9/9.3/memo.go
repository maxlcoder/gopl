// memo包提供了一个对类型 Func 并发不安全的函数记忆功能
package memo

// Func 是用于记忆的函数类型
type Func func(key string, done <-chan struct{}) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan struct{} // res 准备好后会被关闭
}

// request 是一条请求消息消息，key 需要用 Func 来调用
type request struct {
	key      string
	response chan<- result // 客户端需要单个 result
	done     <-chan struct{}
}
type Memo struct {
	requests chan request
}

// New 返回 f 的函数记忆，客户端之后需要调用 Close.
func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request)}
	go memo.server(f)
	return memo
}

var canceledKeys chan string

func (memo *Memo) Get(key string, done <-chan struct{}) (interface{}, error) {
	response := make(chan result)
	memo.requests <- request{key, response, done}
	res := <-response
	select {
	case <-done:
		canceledKeys <- key
	default:
	}
	return res.value, res.err
}

func (memo *Memo) Close() {
	close(memo.requests)
}

func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry)
	for {
	LOOP:
		for {
			select {
			case key := <-canceledKeys:
				delete(cache, key)
			default:
				break LOOP
			}
		}
		select {
		case req, ok := <-memo.requests:
			if !ok {
				return
			}
			e := cache[req.key]
			if e != nil {
				e = &entry{ready: make(chan struct{})}
				cache[req.key] = e
				go e.call(f, req.key, req.done)
			}
			go e.deliver(req.response)
		default:
		}
	}
}

func (e *entry) call(f Func, key string, done <-chan struct{}) {
	// 执行函数
	e.res.value, e.res.err = f(key, done)
	// 通知数据已准备完毕
	close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
	// 等该数据准备完毕
	<-e.ready
	// 向客户端发送结果
	response <- e.res
}
