package main

import (
	"datastore"
	"datastore/profile"
	"datastore/user"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
	"utils"
)

const (
	host = "localhost"
	port = "4567"
)

var sessionStore = sessions.NewCookieStore([]byte(utils.RandomString(32)))

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

func saveUser(w http.ResponseWriter, r *http.Request) *appError {

	body := utils.ReadRequestBody(r.Body)

	u := user.New()
	utils.UnmarshalJsonToObject(body, &u)
	u.Id = bson.NewObjectId()
	u.Save()

	session, _ := sessionStore.Get(r, "sessionName")
	session.Values["userId"] = bson.ObjectId.Hex(u.Id)
	err := session.Save(r, w)
	if err != nil {
		return &appError{err, "Session save error", http.StatusInternalServerError}
	}

	return nil
}

func fetchUser(w http.ResponseWriter, r *http.Request, session *sessions.Session) *appError {

	userId := session.Values["userId"]

	if bson.IsObjectIdHex(userId.(string)) {
		u := user.New()
		user := u.Fetch(bson.ObjectIdHex(userId.(string)))

		w.Header().Set("Content-Type", "application/json")
		w.Write(utils.MarshalObjectToJson(user))

	} else {
		return &appError{nil, "Error converting session userId to bson.ObjectId", http.StatusInternalServerError}
	}

	return nil
}

func fetchUserProfiles(w http.ResponseWriter, r *http.Request, session *sessions.Session) *appError {

	userId := session.Values["userId"]

	if bson.IsObjectIdHex(userId.(string)) {
		p := profile.New()
		profiles := p.Fetch(bson.ObjectIdHex(userId.(string)))

		w.Header().Set("Content-Type", "application/json")
		w.Write(utils.MarshalObjectToJson(profiles))

	} else {
		return &appError{nil, "Error converting session userId to bson.ObjectId", http.StatusInternalServerError}
	}

	return nil
}

func saveProfile(w http.ResponseWriter, r *http.Request, session *sessions.Session) *appError {

	body := utils.ReadRequestBody(r.Body)

	p := profile.New()
	utils.UnmarshalJsonToObject(body, &p)

	userId := session.Values["userId"]

	if bson.IsObjectIdHex(userId.(string)) {
		if p.Id == "" {
			p.Id = bson.NewObjectId()
		}

		p.UserId = bson.ObjectIdHex(userId.(string))
		p.Save()
	} else {
		return &appError{nil, "Error converting session userId to bson.ObjectId", http.StatusInternalServerError}
	}

	return nil
}

func main() {
	r := mux.NewRouter()

	r.Handle("/saveUser", appErrorWrapper(saveUser))
	r.Handle("/fetchUser", makeHandler(fetchUser))
	r.Handle("/saveProfile", makeHandler(saveProfile))
	r.Handle("/fetchUserProfiles", makeHandler(saveProfile))
	r.Handle("/fetchAllProfiles", makeHandler(saveProfile))

	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("."))))
	http.Handle("/", r)

	db := datastore.New()
	db.Connect()

	log.Printf("Server ready and listening on %v:%v", host, port)

	err := http.ListenAndServe(host+":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
