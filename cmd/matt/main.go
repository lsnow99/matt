package main

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

var SecretPages = []string{"HomeView"}
var Password = "secret"

func main() {
	r := mux.NewRouter()

	staticFS := http.Dir("frontend/dist")
	fs := http.FileServer(staticFS)
	r.PathPrefix("/").Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		if isProtected(path) {
			pwCookie, err := r.Cookie("password")
			if err != nil || pwCookie.Value != Password {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
		}

		// static files
		if strings.Contains(path, ".") || path == "/" {
			if path != "/" {
				w.Header().Add("Cache-Control", "max-age=31536000")
			}
			fs.ServeHTTP(w, r)
			return
		}

		// default to serve index.html
		f, _ := staticFS.Open("index.html")
		http.ServeContent(w, r, "index.html", time.Time{}, f)
	}))

	http.Handle("/", r)
	http.ListenAndServe(":"+strconv.Itoa(8000), http.DefaultServeMux)
}

func isProtected(path string) bool {
	for _, page := range SecretPages {
		if strings.Contains(path, page) {
			return true
		}
	}
	return false
}
