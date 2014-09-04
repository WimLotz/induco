package main

import (
	"code.google.com/p/goauth2/oauth"
	"encoding/json"
	"fmt"
	"labix.org/v2/mgo/bson"
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

	body := readRequestBody(r.Body)

	code := string(body)

	accessToken, idToken, err := exchange(code)
	if err != nil {
		return &appError{err, "Error exchanging code for access token", 500}
	}

	gplusID, err := decodeIdToken(idToken)
	if err != nil {
		return &appError{err, "Error decoding ID token", 500}
	}

	storedToken := session.Values["accessToken"]
	storedGPlusID := session.Values["gplusID"]
	if storedToken == nil && storedGPlusID != gplusID {
		session.Values["accessToken"] = accessToken
		session.Values["gplusID"] = gplusID
	}

	repo := createUsersRepo()
	id := repo.fetchUserIdFromGooglePlusId(gplusID)

	if bson.ObjectId.Valid(id) {
		session.Values["userId"] = bson.ObjectId.Hex(id)
	} else {
		newId := bson.NewObjectId()
		repo.createUser(user{Id: newId, GoogleAuthId: gplusID})
		session.Values["userId"] = bson.ObjectId.Hex(newId)
	}

	err = session.Save(r, w)
	if err != nil {
		return &appError{err, "Session save error", http.StatusInternalServerError}
	}

	return nil
}
