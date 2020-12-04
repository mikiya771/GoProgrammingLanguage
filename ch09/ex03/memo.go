package memo

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

type request struct {
	key      string
	response chan<- result
	done     <-chan struct{}
}

type entry struct {
	res   result
	ready chan struct{}
}
type Memo struct {
	requests chan request
}

type Func func(key string, done <-chan struct{}) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request)}
	go memo.server(f)
	return memo
}
func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry)
	for req := range memo.requests {
		e := cache[req.key]
		if e != nil {
			select {
			case <-e.ready:
				if e.res.err == ErrCanceled {
					delete(cache, req.key)
					e = nil
				}
			default:
			}
		}
		if e == nil {
			e = &entry{ready: make(chan struct{})}
			cache[req.key] = e
			go e.call(f, req.key)
		}
		go e.deliver(req.response)
	}
}

func (e *entry) call(f Func, key string, done chan struct{}) {
	e.res.value, e.res.err = f(key, done)
	if e.res.err != nil &&
		strings.Contains(e.res.err.Error(), "request canceled") {
		e.res.err = ErrCanceled
	}
	close(e.ready)
}
func (e *entry) deliver(response chan<- result) {
	<-e.ready
	response <- e.res
}

func (memo *Memo) Get(key string, done chan struct{}) (interface{}, error) {
	response := make(chan result)
	memo.requests <- request{key, response, done}
	res := <-response
	if res.err == ErrCanceled {
		select {
		case <-done:
			return res.value, res.err
		default:
			return memo.Get(key, done)
		}
	}
	return res.value, res.err
}

func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

var ErrCanceled = errors.New("request canceled")
