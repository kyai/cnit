package handler

import (
	"bytes"
	"cnit/crawler"
	"encoding/json"
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	// render html
	t, _ := template.ParseFiles("views/home.html")
	buf := new(bytes.Buffer)
	t.Execute(buf, map[string]interface{}{
		"hello": "world",
	})
	w.Write(buf.Bytes())
}

func Data(w http.ResponseWriter, r *http.Request) {
	job := crawler.Job{}
	// output json
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	b, _ := json.Marshal(job)
	w.Write(b)
}
