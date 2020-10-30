package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"sort"
	"strings"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe("localhost:8000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	td := make([]*Track, len(tracksData))
	copy(td, tracksData)
	value := r.URL.Query().Get("sort")
	f, keys := sortOrders(strings.Split(value, ","))
	if len(f) == 0 {
		printTracksHTML(w, td, keys)
		return
	}

	table := &multikeySortTracks{tracks: td}
	fmt.Println(f)
	for _, sk := range f {
		table.AddSortKey(sk)
	}
	fmt.Println(table.lessFunc)
	sort.Sort(table)
	printTracksHTML(w, td, keys)
}

func createQueryLink(keys []string, name string) template.HTML {
	updatedKeys := make([]string, 0, len(keys))
	for _, key := range keys {
		if strings.ToUpper(key) != strings.ToUpper(name) {
			updatedKeys = append(updatedKeys, key)
		}
	}
	updatedKeys = append(updatedKeys, name)
	queryLink := fmt.Sprintf("<a href=\"?sort=%s\">%s</a>",
		strings.Join(updatedKeys, ","), name)
	return template.HTML(queryLink)
}
func printTracksHTML(out io.Writer, tracks []*Track, keys []string) {
	titleFunc := func() template.HTML { return createQueryLink(keys, "Title") }
	artistFunc := func() template.HTML { return createQueryLink(keys, "Artist") }
	albumFunc := func() template.HTML { return createQueryLink(keys, "Album") }
	yearFunc := func() template.HTML { return createQueryLink(keys, "Year") }
	lengthFunc := func() template.HTML { return createQueryLink(keys, "Length") }

	funcMap := template.FuncMap{
		"title":  titleFunc,
		"artist": artistFunc,
		"album":  albumFunc,
		"year":   yearFunc,
		"length": lengthFunc}

	err := template.Must(template.New("tracktable").
		Funcs(funcMap).
		Parse(`
		<html>
		<head>
		<meta http-equiv="Content-Type" conntent="text/html; charset=utf-8">
		<title>My Tracks</title>
		</head>
		</body>
		<table border="5" rules="all" cellpadding="5">
		<tr style='text-align: left'>
			<th>{{title}}</th>
			<th>{{artist}}</th>
			<th>{{album}}</th>
			<th>{{year}}</th>
			<th>{{length}}</th>
		</tr>
		{{range .Tracks}}
		<tr>
			<td>{{.Title}}</td>
			<td>{{.Artist}}</td>
			<td>{{.Album}}</td>
			<td>{{.Year}}</td>
			<td>{{.Length}}</td>
		</tr>
		{{end}}
		</table>
		</body>
		</html>
	`)).Execute(out, &TracksTable{tracks})
	if err != nil {
		fmt.Printf("%v\n", err)
	}
}

type TracksTable struct {
	Tracks []*Track
}

func sortOrders(keys []string) ([]func(p, q interface{}) bool, []string) {
	updatedKeys := make([]string, 0, len(keys))
	sortOrderFuncs := make([]func(p, q interface{}) bool, 0, len(keys))

	for _, key := range keys {
		key := strings.TrimSpace(key)
		f, ok := sortKeyFuncs[strings.ToUpper(key)]
		if ok {
			updatedKeys = append(updatedKeys, key)
			sortOrderFuncs = append(sortOrderFuncs, f)
		}
	}
	return sortOrderFuncs, updatedKeys
}
