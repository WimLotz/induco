package main

import (
	"code.google.com/p/goauth2/oauth"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"github.com/gorilla/sessions"
	"html/template"
)

type redirect struct {
	Url string `json:"url"`
}

//const profileInfoURL = "https://www.googleapis.com/oauth2/v1/userinfo?alt=json"
const tokenSessionStore = "token-storage"
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
	TokenCache:   oauth.CacheFile("cache.json"),
}

func main() {
	r := mux.NewRouter()

	r.Handle("/", http.RedirectHandler("/static/", 302)).Name("home")
	r.HandleFunc("/googleAuthorise", googleAuthoriseHandler).Methods("POST")
	r.HandleFunc("/oauth2callback", handleOAuth2Callback).Methods("GET")
	r.HandleFunc("/isLoggedIn", isLoggedInHandler).Methods("GET")
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("."))))
	http.Handle("/", r)

	log.Printf("Server ready and listening on %v:%v", host, port)
	http.ListenAndServe(host+":"+port, nil)
}

func isLoggedInHandler(w http.ResponseWriter, r *http.Request){
	session, _ := cookieStore.Get(r, tokenSessionStore)
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
	if err != nil{
		log.Printf("Oauth transport exchange error:%v", err)
		return
	}

	session, err := cookieStore.Get(r, tokenSessionStore)
	if err != nil{
		log.Printf("Session retrieve error:%v", err)
		return
	}

	session.Values["token"] = token.AccessToken
	session.Save(r, w)

	url, err := mux.CurrentRoute(r).Subrouter().Get("home").URL()
	if err != nil {
		log.Fatalf("Unable to get home route. Error:%v", err)
	}

	http.Redirect(w, r, url.String(), 302)
}

var userInfoTemplate = template.Must(template.New("").Parse(`
<html><body>
This app is now authenticated to access your Google user info.  Your details are:<br />
{{.}}
</body></html>
`))
