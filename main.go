package main

import (
	"cnit/handler"
	"net/http"
)

const (
	ServeHost = ":8899"
	RedisHost = ""
)

var routers = map[string]func(http.ResponseWriter, *http.Request){
	"/":     handler.Home,
	"/data": handler.Data,
}

func main() {
	for pattern, handler := range routers {
		http.HandleFunc(pattern, handler)
	}

	http.Handle("/static", http.FileServer(http.Dir("static")))

	err := http.ListenAndServe(ServeHost, nil)
	if err != nil {
		panic(err)
	}
}
