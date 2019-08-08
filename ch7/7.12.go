package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	const templ = `<table>
<tr style='text-align: left'>
	<th>商品</th>
	<th>价格</th>
</tr>
{{range .GOODS}}
<tr>
	<td>{{.NAME}}</td>
	<td>{{.PRICE}}</td>
</tr>
{{end}}
</table>`
	t := template.Must(template.New("teset").Parse(templ))

	var data struct{
		GOODS []struct{
			NAME  string
			PRICE float32
		}
	}
	for item, price := range db {
		var temp struct {
			NAME  string
			PRICE float32
		}

		temp.NAME = item
		temp.PRICE = float32(price)
		data.GOODS = append(data.GOODS, temp)
	}
	if err := t.Execute(w, data); err != nil {
		log.Fatal(err)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
	fmt.Fprintf(w, "%s\n", price)
}

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
