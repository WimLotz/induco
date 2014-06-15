package main

import (
	"code.google.com/p/goauth2/oauth"
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
)

type redirect struct {
	Url string `json:"url"`
}

const profileInfoURL = "https://www.googleapis.com/oauth2/v1/userinfo?alt=json"
const host = "localhost"
const port = "8080"

var googleOauthCfg = &oauth.Config{
	ClientId:     "77809986585-k1i610m84508e10maaiulfdug3p7u4ni.apps.googleusercontent.com",
	ClientSecret: "RA6UgSeyadzYnrazbHT-86J8",
	AuthURL:      "https://accounts.google.com/o/oauth2/auth",
	TokenURL:     "https://accounts.google.com/o/oauth2/token",
	RedirectURL:  "http://localhost:8080/oauth2callback",
	Scope:        "https://www.googleapis.com/auth/userinfo.profile",
	TokenCache:   oauth.CacheFile("cache.json"),
}

func main() {
	r := mux.NewRouter()

	r.Handle("/", http.RedirectHandler("/static/", 302))
	r.HandleFunc("/googleAuthorise", googleAuthoriseHandler).Methods("POST")
	r.HandleFunc("/oauth2callback", handleOAuth2Callback)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("."))))
	http.Handle("/", r)

	log.Printf("Server ready and listening on %v:%v", host, port)
	http.ListenAndServe(host+":"+port, nil)
}

func googleAuthoriseHandler(w http.ResponseWriter, r *http.Request) {
	url := googleOauthCfg.AuthCodeURL("")
	bytes, err := json.Marshal(redirect{url})

	if err != nil {
		log.Printf("Json Marshalling error:%v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}

func handleOAuth2Callback(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("code")
	transport := &oauth.Transport{Config: googleOauthCfg}
	token, err := transport.Exchange(code)
	transport.Token = token

	result, err := transport.Client().Get(profileInfoURL)
	if err != nil {
		log.Fatal("Oauth transport Get error:", err)
	}

	defer result.Body.Close()

	log.Printf("code:%v", code)     //todo remove debugging code
	io.Copy(os.Stdout, result.Body) //todo remote debugging code
}
