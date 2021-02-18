package main

import (
	"fmt"
	"net/http"
)

//func handler(writer http.ResponseWriter, request *http.Request)  {
//	fmt.Fprintf(writer, "Helle World, %s!", request.URL.Path[1:])
//}

func main() {
	//http.HandleFunc("/", handler)
	//http.ListenAndServe(":8080", nil)

	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("/public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", index)
	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
