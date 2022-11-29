package serve

import (
	"faya/cronlist"
	"faya/db"
	"faya/list"
	"faya/serve/servestrategy"
	"fmt"
	"net/http"
	"strings"
	"time"
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
	fmt.Println(time.Now(), r.URL.Path)
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
    //if r.URL.Path == "/ss1" {
    if strings.HasPrefix(r.URL.Path, "/ss1") {
		if r.URL.Path == "/ss1list" {
			s := &servestrategy.ServeStrategy1{}
			b := s.GetKeyList()
			w.Write(b)
			return
		}


		date := r.URL.Path[4:]
		if len(date) < 2 {
			date = ""
		}else {
			if date[0] == '/' {
				date = date[1:]
			}
		}
		s := &servestrategy.ServeStrategy1{}
		b := s.GetCached(date)
		w.Write(b)
		//w.WriteHeader(200)
		return
	}
	//jobs
	if strings.HasPrefix(r.URL.Path, "/job/") {
		jobname := r.URL.Path[5:]
		var jt *cronlist.JobListUnit = nil
		for _, ju := range cronlist.JobList {
			if ju.Name == jobname {
				jt = &ju
				break
			}
		}
		fmt.Println("I got job :", jobname, jt)
		if jt != nil {
			go jt.Func()
			w.Write([]byte("{\"success\":true, \"msg\":\"job create success\"}"))
			return
		}
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
