package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type XkcdResponse struct {
	ID         int    `db:"id"`
	Url        string `db:"url"`
	Title      string `db:"title"`
	Transcript string `db:"transcript"`
}
type HtResponse struct {
	Num        int    `json:"num"`
	Link       string `json:"link"`
	Year       string `json:"year"`
	News       string `json:"news"`
	SafeTitle  string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	Title      string `json:"title"`
	Day        string `json:"day"`
}

var sqldir string
var sqlpath string

func init() {
	home, err := os.UserHomeDir()
	if err != nil {
		panic("please set $HOME")
	}
	sqldir = filepath.Join(home, ".cache", "mikya771")
	sqlpath = filepath.Join(sqldir, "xkcd_sqlite.db")
}

func main() {
	if len(os.Args) < 1 {
		return
	}
	sqlDB, err := loading()
	if err != nil {
		log.Println(err)
		return
	}
	defer sqlDB.Close()
	search := flag.NewFlagSet("search", flag.ExitOnError)
	clone := flag.NewFlagSet("clone", flag.ExitOnError)
	switch os.Args[1] {
	case "search":
		str := search.String("filter", "", "filter word")
		search.Parse(os.Args[2:])
		retArr, err := searchTitle(sqlDB, *str)
		if err != nil {
			log.Printf("Title Searching Error: %v", err)
			return
		}
		for _, res := range retArr {
			fmt.Println(res)
		}
	case "clone":
		str := clone.String("url", "", "target url")
		fmt.Print(*str)
		clone.Parse(os.Args[2:])
		title, err := cloneUrl(sqlDB, *str)
		fmt.Print(title)
		if err != nil {
			log.Printf("Clone Error: %v", err)
			return
		}
		fmt.Printf("Clone from %s", title)
	}
}

func loading() (*sqlx.DB, error) {
	log.Println("Loading sqlite-database.db...")
	fmt.Println(sqlpath)
	if _, err := os.Stat(sqlpath); err != nil {
		if !os.IsNotExist(err) {
			fmt.Println(err)
			return nil, err
		} else {
			err = initalize()
			if err != nil {
				return nil, err
			}
		}
	}
	sqliteDatabase, err := sqlx.Open("sqlite3", sqlpath)
	if err != nil {
		return nil, err
	}
	cmd := `CREATE TABLE IF NOT EXISTS comic(
	id integer NOT NULL PRIMARY KEY AUTOINCREMENT,	
	url TEXT,    
	transcript TEXT,
	title TEXT);`
	stmt, err := sqliteDatabase.Prepare(cmd)
	if err != nil {
		sqliteDatabase.Close()
		return nil, err
	}
	stmt.Exec()

	return sqliteDatabase, nil
}
func initalize() error {
	fmt.Print(sqldir)
	if err := os.MkdirAll(sqldir, 0777); err != nil {
		return err
	}
	os.Create(sqlpath)
	if file, err := os.Create(sqlpath); err != nil {
		return err
	} else {
		defer file.Close()
	}
	return nil
}

func searchTitle(sql *sqlx.DB, word string) ([]XkcdResponse, error) {
	q := `SELECT id, title, url, transcript FROM comic WHERE title LIKE ? ;`
	var res []XkcdResponse
	fmt.Println(word)
	if err := sql.Select(&res, q, "%"+word+"%"); err != nil {
		return nil, err
	}
	return res, nil
}
func cloneUrl(sql *sqlx.DB, url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	var res HtResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return "", err
	}
	title := res.Title
	transcript := res.Transcript

	q := `INSERT INTO comic(title, url, transcript) VALUES (?,?,?)`
	stmt, err := sql.Prepare(q)
	if err != nil {
		return "", err
	}
	_, err = stmt.Exec(title, url, transcript)
	if err != nil {
		return "", err
	}
	return title, nil
}
