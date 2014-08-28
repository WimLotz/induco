package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"io/ioutil"
	"labix.org/v2/mgo/bson"
	"log"
	"net/http"
)

const (
	host = "localhost"
	port = "4567"
)

var sessionStore = sessions.NewCookieStore([]byte(randomString(32)))

type (
	appError struct {
		Error   error
		Message string
		Code    int
	}
	appHandler func(http.ResponseWriter, *http.Request) *appError
)

func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if e := fn(w, r); e != nil { // e is *appError, not os.Error.
		log.Println(e.Error)
		http.Error(w, e.Message, e.Code)
	}
}

func saveProfile(w http.ResponseWriter, r *http.Request) *appError {

	session, err := sessionStore.Get(r, "sessionName")
	if err != nil {
		log.Printf("unable to retieve sessoion: %v", err)
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("error occured reading from requst body:%v", err)
	}

	var p person
	err = json.Unmarshal(body, &p)
	if err != nil {
		log.Printf("json unmarshalling error:%v", err)
	}

	docId := session.Values["docId"]
	repo := createPeopleRepo()
	if bson.IsObjectIdHex(docId.(string)) {
		log.Printf("valid: %v", bson.ObjectIdHex(docId.(string)))
		repo.updatePerson(p, bson.ObjectIdHex(docId.(string)))
	} else {
		log.Printf("this sucks: %v", docId)
	}

	return nil
}

func main() {
	r := mux.NewRouter()

	r.Handle("/googleConnect", appHandler(googleAuthConnect))
	r.Handle("/saveProfile", appHandler(saveProfile))

	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("."))))
	http.Handle("/", r)

	log.Printf("Server ready and listening on %v:%v", host, port)

	err := http.ListenAndServe(host+":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
