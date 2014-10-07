package main

import (
	"datastore"
	"datastore/company"
	"datastore/person"
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

func savePersonProfile(w http.ResponseWriter, r *http.Request, session *sessions.Session) *appError {

	body := utils.ReadRequestBody(r.Body)

	p := person.New()
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

func saveCompanyProfile(w http.ResponseWriter, r *http.Request, session *sessions.Session) *appError {

	body := utils.ReadRequestBody(r.Body)

	c := company.New()
	utils.UnmarshalJsonToObject(body, &c)

	userId := session.Values["userId"]

	if bson.IsObjectIdHex(userId.(string)) {
		if c.Id == "" {
			c.Id = bson.NewObjectId()
		}

		c.UserId = bson.ObjectIdHex(userId.(string))
		c.Save()
	} else {
		return &appError{nil, "Error converting session userId to bson.ObjectId", http.StatusInternalServerError}
	}

	return nil
}

func fetchPersonProfiles(w http.ResponseWriter, r *http.Request, session *sessions.Session) *appError {

	userId := session.Values["userId"]
	if bson.IsObjectIdHex(userId.(string)) {
		repo := person.CreatePeopleRepo()
		personProfiles := repo.FetchPersonProfiles(bson.ObjectIdHex(userId.(string)))

		w.Header().Set("Content-Type", "application/json")
		w.Write(utils.MarshalObjectToJson(personProfiles))

	} else {
		return &appError{nil, "Error converting session userId to bson.ObjectId", http.StatusInternalServerError}
	}

	return nil
}

func fetchCompanyProfiles(w http.ResponseWriter, r *http.Request, session *sessions.Session) *appError {

	userId := session.Values["userId"]
	if bson.IsObjectIdHex(userId.(string)) {
		repo := company.CreateCompaniesRepo()
		companyProfiles := repo.FetchCompanyProfiles(bson.ObjectIdHex(userId.(string)))

		w.Header().Set("Content-Type", "application/json")
		w.Write(utils.MarshalObjectToJson(companyProfiles))

	} else {
		return &appError{nil, "Error converting session userId to bson.ObjectId", http.StatusInternalServerError}
	}

	return nil
}

func fetchAllProfiles(w http.ResponseWriter, r *http.Request, session *sessions.Session) *appError {

	personRepo := person.CreatePeopleRepo()
	peopleProfiles := personRepo.All()

	companyRepo := company.CreateCompaniesRepo()
	companyProfiles := companyRepo.All()

	allProfiles := make([]interface{}, 50)

	for _, p := range peopleProfiles {
		log.Printf("person %v", p)
	}

	for _, c := range companyProfiles {
		log.Printf("company %v", c)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(utils.MarshalObjectToJson(allProfiles))

	return nil
}

func main() {
	r := mux.NewRouter()

	r.Handle("/googleConnect", appErrorWrapper(googleAuthConnect))
	r.Handle("/googleDisconnect", makeHandler(googleDisconnect))
	r.Handle("/fetchPersonProfiles", makeHandler(fetchPersonProfiles))
	r.Handle("/fetchCompanyProfiles", makeHandler(fetchCompanyProfiles))
	r.Handle("/savePersonProfile", makeHandler(savePersonProfile))
	r.Handle("/saveCompanyProfile", makeHandler(saveCompanyProfile))
	r.Handle("/fetchAllProfiles", makeHandler(fetchAllProfiles))

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
