package serve

import (
	"fmt"
	"net/http"
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
    if r.URL.Path == "/ss1" {
		s := &ServeStrategy1{}
		b := s.GetCached()
		w.Write(b)
		//w.WriteHeader(200)
		return
	}
    http.NotFound(w, r)
    return
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello MyMuxRoute!")
}
