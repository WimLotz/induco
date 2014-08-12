package main

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"

	"log"
	"net/http"
)

const host = "localhost"
const port = "4567"

var sessionStore = sessions.NewCookieStore([]byte(randomString(32)))

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/connect", googleAuthConnect)
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("."))))
	http.Handle("/", r)

	//r.PathPrefix("/static/").Handler(http.StripPrefix("/static",http.FileServer(http.Dir("../static"))))
	//r.HandleFunc("/", serveHello)

	//http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("../static"))))
	//http.HandleFunc("/", serveHello)

	//http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("."))))
	//http.HandleFunc("/", serveHello)

	log.Printf("Server ready and listening on %v:%v", host, port)

	err := http.ListenAndServe(host+":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
