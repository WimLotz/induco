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
)

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
		p.Id = bson.NewObjectId()
		p.UserId = bson.ObjectIdHex(userId.(string))
		p.save()
	} else {
		return &appError{nil, "Error converting session userId to bson.ObjectId", 500}
	}

	saveSession(w, r, session)

	return nil
}

//func saveCompanyProfile(w http.ResponseWriter, r *http.Request) *appError {

//	session, err := sessionStore.Get(r, "sessionName")
//	if err != nil {
//		log.Printf("unable to retieve session: %v", err)
//		http.Error(w, err.Error(), 500)
//		return nil
//	}

//	body, err := ioutil.ReadAll(r.Body)
//	if err != nil {
//		log.Printf("error occured reading from request body: %v", err)
//		http.Error(w, err.Error(), 500)
//		return nil
//	}

//	var c company
//	err = json.Unmarshal(body, &c)
//	if err != nil {
//		log.Printf("json unmarshalling error: %v", err)
//		http.Error(w, err.Error(), 500)
//		return nil
//	}

//	docId := session.Values["docId"]
//	repo := createCompaniesRepo()
//	if bson.IsObjectIdHex(docId.(string)) {
//		c.Id = bson.ObjectIdHex(docId.(string))
//		repo.saveCompany(&c)
//	} else {
//		log.Printf("error converting session docId to bson.ObjectId")
//	}

//	return nil
//}

func fetchPersonProfiles(w http.ResponseWriter, r *http.Request, session *sessions.Session) *appError {

	userId := session.Values["userId"]
	if bson.IsObjectIdHex(userId.(string)) {
		repo := createPeopleRepo()
		personProfile := repo.fetchPersonProfiles(bson.ObjectIdHex(userId.(string)))

		w.Header().Set("Content-Type", "application/json")
		w.Write(marshalObjectToJson(personProfile))

	} else {
		return &appError{nil, "Error converting session userId to bson.ObjectId", 500}
	}

	return nil
}

//func fetchCompanyProfile(w http.ResponseWriter, r *http.Request) *appError {

//	session, err := sessionStore.Get(r, "sessionName")
//	if err != nil {
//		log.Printf("unable to retieve session: %v", err)
//	}

//	docId := session.Values["docId"]
//	repo := createCompaniesRepo()
//	if bson.IsObjectIdHex(docId.(string)) {
//		companyProfile := repo.fetchCompanyProfile(bson.ObjectIdHex(docId.(string)))
//		jsonData, err := json.Marshal(companyProfile)
//		if err != nil {
//			http.Error(w, err.Error(), 500)
//			return nil
//		}

//		w.Header().Set("Content-Type", "application/json")
//		w.Write(jsonData)

//	} else {
//		log.Printf("error converting session docId to bson.ObjectId")
//	}

//	return nil
//}

func main() {
	r := mux.NewRouter()

	//r.Handle("/googleConnect", appHandler(googleAuthConnect))
	r.Handle("/fetchPersonProfiles", makeHandler(fetchPersonProfiles))
	//r.Handle("/fetchCompanyProfile", makeHandler(fetchCompanyProfile))
	r.Handle("/savePersonProfile", makeHandler(savePersonProfile))
	//r.Handle("/saveCompanyProfile", makeHandler(saveCompanyProfile))

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
