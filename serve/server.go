package serve

import (
	"faya/db"
	"faya/list"
	"fmt"
	"net/http"
	"strings"
)


type StandStill struct {
}

func enableCors(w *http.ResponseWriter) {
    (*w).Header().Set("Access-Control-Allow-Origin", "*")
    (*w).Header().Set("Access-Control-Allow-Origin", "*")
    (*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    (*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func (p *StandStill) ServeHTTP(w http.ResponseWriter, r *http.Request ){
	enableCors(&w)
    if r.URL.Path == "/" {
        sayhelloName(w, r)
        return
    }
    if r.URL.Path == "/isopen" {
		isOpening := list.IsOpeningTime()
		var ret []byte
		if isOpening{
			ret = []byte("{\"open\":true}")
		}else {
			ret = []byte("{\"open\":false}")
		}
		w.Write(ret)
		return
	}
    if r.URL.Path == "/ss1" {
		s := &ServeStrategy1{}
		b := s.GetCached()
		w.Write(b)
		//w.WriteHeader(200)
		return
	}

	//read from redis
	if strings.HasPrefix(r.URL.Path, "/db/") {
		path := r.URL.Path[4:]
		valid := true
		for _, c := range path {
			if c >= '0' && c <= '9' ||
			c >= 'A' && c <= 'Z' ||
			c >= 'a' && c <= 'z' ||
			c == '_' ||
			c == '-' {
			}else {
				valid = false
				break
			}
		}
		fmt.Println("I got for redis key:", path, valid)
		if valid {
			contentStr, had := db.SimpleGet(path)
			if had {
				w.Write([]byte(contentStr))
				return
			}
		}
	}

	
    http.NotFound(w, r)
    return
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello MyMuxRoute!")
}
