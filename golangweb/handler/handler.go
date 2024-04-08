package handler

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	// w.Write([]byte("selamat datang"))
tmpl, err := template.ParseFiles(path.Join("views", "index.html"))
	if err != nil   {
		log.Println(err)
		http.Error(w, "Sedang ada masalah cek kembali", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"title": "Belajar Golang",
		"content": "im learning golang web by my self",
	}

	err = tmpl.Execute(w, data)
if err != nil {
	log.Println(err)
	http.Error(w,"Sedang ada masalah cek kembali", http.StatusInternalServerError)
	return
}

	// w.Write(([]byte("Home")))
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world saya sedang belajar golang web"))
}
func MarioHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("mantep sudah bisa"))
}
func ProductHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idNumb, err := strconv.Atoi(id)
	if err != nil || idNumb < 1 {
		http.NotFound(w, r)
		return
	}
	// w.Write([]byte("product page"))

	fmt.Fprintf(w, "Product page : %d", idNumb)
}