package main

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
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
	appErrorWrapper func(http.ResponseWriter, *http.Request) *appError
)

func (fn appErrorWrapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if e := fn(w, r); e != nil {
		log.Println(e.Error)
		http.Error(w, e.Message, e.Code)
	}
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, *sessions.Session) *appError) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		session, err := sessionStore.Get(r, "sessionName")
		if err != nil {
			log.Printf("Error occured retrieving session: %v\n", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		if appErr := fn(w, r, session); appErr != nil {
			log.Println(appErr.Error)
			http.Error(w, appErr.Message, appErr.Code)
		}
	}
}

func savePersonProfile(w http.ResponseWriter, r *http.Request, session *sessions.Session) *appError {

	body := readRequestBody(r.Body)

	var p person
	unmarshalJsonToObject(body, &p)

	userId := session.Values["userId"]

	if bson.IsObjectIdHex(userId.(string)) {
		if p.Id == "" {
			p.Id = bson.NewObjectId()
		}

		p.UserId = bson.ObjectIdHex(userId.(string))
		p.save()
	} else {
		return &appError{nil, "Error converting session userId to bson.ObjectId", http.StatusInternalServerError}
	}

	saveSession(w, r, session)

	return nil
}

func saveCompanyProfile(w http.ResponseWriter, r *http.Request, session *sessions.Session) *appError {

	body := readRequestBody(r.Body)

	var c company
	unmarshalJsonToObject(body, &c)

	userId := session.Values["userId"]

	if bson.IsObjectIdHex(userId.(string)) {
		if c.Id == "" {
			c.Id = bson.NewObjectId()
		}

		c.UserId = bson.ObjectIdHex(userId.(string))
		c.save()
	} else {
		return &appError{nil, "Error converting session userId to bson.ObjectId", http.StatusInternalServerError}
	}

	return nil
}

func fetchPersonProfiles(w http.ResponseWriter, r *http.Request, session *sessions.Session) *appError {

	userId := session.Values["userId"]
	if bson.IsObjectIdHex(userId.(string)) {
		repo := createPeopleRepo()
		personProfiles := repo.fetchPersonProfiles(bson.ObjectIdHex(userId.(string)))

		w.Header().Set("Content-Type", "application/json")
		w.Write(marshalObjectToJson(personProfiles))

	} else {
		return &appError{nil, "Error converting session userId to bson.ObjectId", http.StatusInternalServerError}
	}

	return nil
}

func fetchCompanyProfiles(w http.ResponseWriter, r *http.Request, session *sessions.Session) *appError {

	userId := session.Values["userId"]
	if bson.IsObjectIdHex(userId.(string)) {
		repo := createCompaniesRepo()
		companyProfiles := repo.fetchCompanyProfiles(bson.ObjectIdHex(userId.(string)))

		w.Header().Set("Content-Type", "application/json")
		w.Write(marshalObjectToJson(companyProfiles))

	} else {
		return &appError{nil, "Error converting session userId to bson.ObjectId", http.StatusInternalServerError}
	}

	return nil
}

func main() {
	r := mux.NewRouter()

	r.Handle("/googleConnect", appErrorWrapper(googleAuthConnect))
	r.Handle("/fetchPersonProfiles", makeHandler(fetchPersonProfiles))
	r.Handle("/fetchCompanyProfiles", makeHandler(fetchCompanyProfiles))
	r.Handle("/savePersonProfile", makeHandler(savePersonProfile))
	r.Handle("/saveCompanyProfile", makeHandler(saveCompanyProfile))

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
