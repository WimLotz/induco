package main

import (
	"code.google.com/p/goauth2/oauth"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"log"
	"net/http"
)

type redirect struct {
	Url string `json:"url"`
}

type user struct {
	IsLoggedIn bool `json:"isLoggedIn"`
}

//const profileInfoURL = "https://www.googleapis.com/oauth2/v1/userinfo?alt=json"
const tokenSessionStore = "token-storage-session"
const tokenForTokenSessionStore = "token"
const host = "localhost"
const port = "8080"

var cookieStore = sessions.NewCookieStore([]byte("something-very-secret"))

var googleOauthCfg = &oauth.Config{
	ClientId:     "77809986585-k1i610m84508e10maaiulfdug3p7u4ni.apps.googleusercontent.com",
	ClientSecret: "RA6UgSeyadzYnrazbHT-86J8",
	AuthURL:      "https://accounts.google.com/o/oauth2/auth",
	TokenURL:     "https://accounts.google.com/o/oauth2/token",
	RedirectURL:  "http://localhost:8080/oauth2callback",
	Scope:        "https://www.googleapis.com/auth/userinfo.profile",
}

func main() {
	r := mux.NewRouter()

	r.Handle("/", http.RedirectHandler("/static/", 302)).Name("home")
	r.HandleFunc("/googleAuthorise", googleAuthoriseHandler).Methods("POST")
	r.HandleFunc("/oauth2callback", handleOAuth2Callback)
	r.HandleFunc("/isLoggedIn", isLoggedInHandler).Methods("GET")
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("."))))
	http.Handle("/", r)

	log.Printf("Server ready and listening on %v:%v", host, port)
	http.ListenAndServe(host+":"+port, nil)
}

func isLoggedInHandler(w http.ResponseWriter, r *http.Request) {
	session := getSession(r, tokenSessionStore)
	currentUser := user{}

	if session.Values[tokenForTokenSessionStore] != nil {
		currentUser.IsLoggedIn = true
	} else {
		currentUser.IsLoggedIn = false
	}

	bytes := marshalObjectToJson(currentUser)
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}

func googleAuthoriseHandler(w http.ResponseWriter, r *http.Request) {
	url := googleOauthCfg.AuthCodeURL("")

	bytes := marshalObjectToJson(redirect{url})
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}

func handleOAuth2Callback(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("code")
	transport := &oauth.Transport{Config: googleOauthCfg}

	token, err := transport.Exchange(code)
	if err != nil {
		log.Printf("Oauth transport exchange error:%v", err)
		return
	}

	session := getSession(r, tokenSessionStore)

	session.Values[tokenForTokenSessionStore] = token.AccessToken
	session.Save(r, w)

	url, err := mux.CurrentRoute(r).Subrouter().Get("home").URL()
	if err != nil {
		log.Fatalf("Unable to get home route. Error:%v", err)
	}

	http.Redirect(w, r, url.String(), 302)
}

func getSession(r *http.Request, sessionName string) *sessions.Session {
	session, err := cookieStore.Get(r, sessionName)
	if err != nil {
		log.Printf("Session retrieve error:%v", err)
		return nil
	}
	return session
}

func marshalObjectToJson(obj interface{}) []byte {
	bytes, err := json.Marshal(obj)
	if err != nil {
		log.Printf("Json Marshalling error:%v", err)
		return nil
	}
	return bytes
}
