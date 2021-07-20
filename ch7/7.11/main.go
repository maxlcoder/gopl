package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type dollars float32

func (d dollars) String() string  {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request)  {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request)  {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
	fmt.Fprintf(w, "%s\n", price)
}

func (db database) update(w http.ResponseWriter, req *http.Request)  {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
	paramPrice := req.URL.Query().Get("price")
	if  paramPrice == "" {
		w.WriteHeader(http.StatusBadRequest) // 404
		fmt.Fprintf(w, "价格缺失")
		return
	}

	// 更新价格
	val, err := strconv.ParseFloat(paramPrice, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // 404
		fmt.Fprintf(w, "参数错误")
		return
	}
	db[item] = dollars(val)
	price = db[item]

	fmt.Fprintf(w, "新价格为： %s\n", price)
}


func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/update", db.update)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}


