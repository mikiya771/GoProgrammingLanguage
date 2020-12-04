package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

var db database

func main() {
	db = database{"shoes": 50, "socks": 5}
	mux := http.NewServeMux()
	mux.Handle("/list", http.HandlerFunc(db.list))
	mux.Handle("/update", http.HandlerFunc(db.update))
	mux.Handle("/delete", http.HandlerFunc(db.delete))
	mux.Handle("/price", http.HandlerFunc(db.price))
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	t := template.Must(template.ParseFiles("index.html"))
	Key := req.URL.Query().Get("select")
	Value, ok := db[Key]
	if !ok {
		TplValue := map[string]interface{}{
			"Values": db,
			"Key":    "",
			"Value":  "",
		}
		fmt.Println("empty update form")
		t.Execute(w, TplValue)
	} else {
		TplValue := map[string]interface{}{
			"Values": db,
			"Key":    Key,
			"Value":  Value,
		}
		fmt.Printf("prepared update form %s, %s \n", Key, Value)
		t.Execute(w, TplValue)
	}
}
func (db database) delete(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		fmt.Println("Valid Method")
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "invalid")
		return
	}
	e := req.ParseForm()
	if e != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "invalid input")
		return
	}
	name := req.Form.Get("name")
	delete(db, name)
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("location", "/list")
	w.WriteHeader(http.StatusMovedPermanently)

}
func (db database) update(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		fmt.Println("Valid Method")
	default:
		w.WriteHeader(http.StatusNotFound)
		return
	}
	e := req.ParseForm()
	if e != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "invalid input")
		return
	}
	name := req.Form.Get("name")
	value, err := strconv.ParseFloat(req.Form.Get("value"), 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s is invalid input: value shoud be float number", req.Form.Get("value"))
		return
	}
	db[name] = dollars(value)
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("location", "/list")
	w.WriteHeader(http.StatusMovedPermanently)
}
func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}
