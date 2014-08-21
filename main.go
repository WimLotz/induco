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

type appError struct {
	Error   error
	Message string
	Code    int
}

type appHandler func(http.ResponseWriter, *http.Request) *appError

func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if e := fn(w, r); e != nil { // e is *appError, not os.Error.
		log.Println(e.Error)
		http.Error(w, e.Message, e.Code)
	}
}

func dashboard(w http.ResponseWriter, r *http.Request) *appError {

	session, err := sessionStore.Get(r, "sessionName")
	if err != nil {
		log.Println("error fetching session1:", err)
		return &appError{err, "Error fetching session", 500}
	}
	log.Println("session:", session.Values["state"].(string))

	//if r.FormValue("state") != session.Values["state"].(string) {
	//	m := "Invalid state parameter"
	//	return &appError{errors.New(m), m, 401}
	//}

	return nil
}

func main() {
	r := mux.NewRouter()

	r.Handle("/googleConnect", appHandler(googleAuthConnect))
	r.Handle("/dashboard", appHandler(dashboard))

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
