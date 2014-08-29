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
		p.Id = bson.ObjectIdHex(docId.(string))
		repo.updatePerson(p)
	} else {
		log.Printf("error converting session docId to bson.ObjectId")
	}

	return nil
}

func fetchProfile(w http.ResponseWriter, r *http.Request) *appError {

	session, err := sessionStore.Get(r, "sessionName")
	if err != nil {
		log.Printf("unable to retieve sessoion: %v", err)
	}

	docId := session.Values["docId"]
	repo := createPeopleRepo()
	if bson.IsObjectIdHex(docId.(string)) {
		personProfile := repo.fetchProfile(bson.ObjectIdHex(docId.(string)))
		jsonData, err := json.Marshal(personProfile)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return nil
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)

	} else {
		log.Printf("error converting session docId to bson.ObjectId")
	}

	return nil
}

func main() {
	r := mux.NewRouter()

	r.Handle("/googleConnect", appHandler(googleAuthConnect))
	r.Handle("/fetchProfile", appHandler(fetchProfile))
	r.Handle("/saveProfile", appHandler(saveProfile))

	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("."))))
	http.Handle("/", r)

	var db dataBase
	db.connect()

	log.Printf("Server ready and listening on %v:%v", host, port)

	err := http.ListenAndServe(host+":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
