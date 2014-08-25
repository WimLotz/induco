package main

import (
	"code.google.com/p/goauth2/oauth"
	"encoding/json"
	//"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	clientID     = "907945139632-72um3uat0epi2o59fb4ci0b2a2sls7o9.apps.googleusercontent.com"
	clientSecret = "Zru8XKpJmjuP8jxRuFNbsyy_"
)

var config = &oauth.Config{
	ClientId:     clientID,
	ClientSecret: clientSecret,
	// Scope determines which API calls you are authorized to make
	Scope:    "https://www.googleapis.com/auth/plus.login",
	AuthURL:  "https://accounts.google.com/o/oauth2/auth",
	TokenURL: "https://accounts.google.com/o/oauth2/token",
	// Use "postmessage" for the code-flow for server side apps
	RedirectURL: "postmessage",
}

type (
	Token struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		ExpiresIn   int    `json:"expires_in"`
		IdToken     string `json:"id_token"`
	}
	// ClaimSet represents an IdToken response.
	ClaimSet struct {
		Sub string
	}
)

// exchange takes an authentication code and exchanges it with the OAuth
// endpoint for a Google API bearer token and a Google+ ID
func exchange(code string) (accessToken string, idToken string, err error) {
	// Exchange the authorization code for a credentials object via a POST request
	addr := "https://accounts.google.com/o/oauth2/token"
	values := url.Values{
		"Content-Type":  {"application/x-www-form-urlencoded"},
		"code":          {code},
		"client_id":     {clientID},
		"client_secret": {clientSecret},
		"redirect_uri":  {config.RedirectURL},
		"grant_type":    {"authorization_code"},
	}
	resp, err := http.PostForm(addr, values)
	if err != nil {
		return "", "", fmt.Errorf("Exchanging code: %v", err)
	}
	defer resp.Body.Close()

	// Decode the response body into a token object
	var token Token
	err = json.NewDecoder(resp.Body).Decode(&token)
	if err != nil {
		return "", "", fmt.Errorf("Decoding access token: %v", err)
	}

	return token.AccessToken, token.IdToken, nil
}

// decodeIdToken takes an ID Token and decodes it to fetch the Google+ ID within
func decodeIdToken(idToken string) (gplusID string, err error) {
	// An ID token is a cryptographically-signed JSON object encoded in base 64.
	// Normally, it is critical that you validate an ID token before you use it,
	// but since you are communicating directly with Google over an
	// intermediary-free HTTPS channel and using your Client Secret to
	// authenticate yourself to Google, you can be confident that the token you
	// receive really comes from Google and is valid. If your server passes the ID
	// token to other components of your app, it is extremely important that the
	// other components validate the token before using it.
	var set ClaimSet
	if idToken != "" {
		// Check that the padding is correct for a base64decode
		parts := strings.Split(idToken, ".")
		if len(parts) < 2 {
			return "", fmt.Errorf("Malformed ID token")
		}
		// Decode the ID token
		b, err := base64Decode(parts[1])
		if err != nil {
			return "", fmt.Errorf("Malformed ID token: %v", err)
		}
		err = json.Unmarshal(b, &set)
		if err != nil {
			return "", fmt.Errorf("Malformed ID token: %v", err)
		}
	}
	return set.Sub, nil
}

func googleAuthConnect(w http.ResponseWriter, r *http.Request) *appError {

	session, _ := sessionStore.Get(r, "sessionName")

	// todo request forgery protection
	// set state
	//if session.Values["state"] == nil {
	//	state := randomString(64)
	//	session.Values["state"] = state
	//}
	// get state and check against ??
	//if r.FormValue("state") != session.Values["state"].(string) {
	//	m := "Invalid state parameter"
	//	return &appError{errors.New(m), m, 401}
	//}

	x, err := ioutil.ReadAll(r.Body)
	fmt.Printf("some id: %v\n", string(x))
	if err != nil {
		return &appError{err, "Error reading code in request body", 500}
	}
	code := string(x)

	accessToken, idToken, err := exchange(code)
	if err != nil {
		return &appError{err, "Error exchanging code for access token", 500}
	}

	gplusID, err := decodeIdToken(idToken)
	if err != nil {
		return &appError{err, "Error decoding ID token", 500}
	}

	//storedToken := session.Values["accessToken"]
	//storedGPlusID := session.Values["gplusID"]
	//if storedToken != nil && storedGPlusID == gplusID {
	//	m := "Current user already connected"
	//	return &appError{errors.New(m), m, 200}
	//}

	session.Values["accessToken"] = accessToken
	session.Values["gplusID"] = gplusID
	fmt.Printf("gplusID: %v\n", string(gplusID))

	session.Save(r, w)

	repo := createPeopleRepo()
	repo.createPerson(person{Id: gplusID, FirstName: "Kinjal", Surname: "Vithal"})

	//newUser := user{Id: gplusID, FirstName: "Kinjal", Surname: "Vithal"}
	//newUser.createNewUser()

	return nil
}

// disconnect revokes the current user's token and resets their session
//func disconnect(w http.ResponseWriter, r *http.Request) *appError {
//	// Only disconnect a connected user
//	session, err := store.Get(r, "sessionName")
//	if err != nil {
//		log.Println("error fetching session:", err)
//		return &appError{err, "Error fetching session", 500}
//	}
//	token := session.Values["accessToken"]
//	if token == nil {
//		m := "Current user not connected"
//		return &appError{errors.New(m), m, 401}
//	}

//	// Execute HTTP GET request to revoke current token
//	url := "https://accounts.google.com/o/oauth2/revoke?token=" + token.(string)
//	resp, err := http.Get(url)
//	if err != nil {
//		m := "Failed to revoke token for a given user"
//		return &appError{errors.New(m), m, 400}
//	}
//	defer resp.Body.Close()

//	// Reset the user's session
//	session.Values["accessToken"] = nil
//	session.Save(r, w)
//	return nil
//}
