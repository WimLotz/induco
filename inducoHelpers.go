package main

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"github.com/gorilla/sessions"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func saveSession(w http.ResponseWriter, r *http.Request, session *sessions.Session) {
	err := session.Save(r, w)
	if err != nil {
		log.Printf("Session save error: %v\n", err)
	}
}

func readRequestBody(r io.Reader) []byte {
	body, err := ioutil.ReadAll(r)
	if err != nil {
		log.Printf("Error occured reading from request body: %v\n", err)
	}
	return body
}

func marshalObjectToJson(obj interface{}) []byte {
	bytes, err := json.Marshal(obj)
	if err != nil {
		log.Printf("Json Marshalling error: %v\n", err)
		return nil
	}
	return bytes
}

func unmarshalJsonToObject(data []byte, v interface{}) {
	err := json.Unmarshal(data, &v)
	if err != nil {
		log.Printf("Json Unmarshalling error: %v\n", err)
	}
}

func randomString(length int) (str string) {
	b := make([]byte, length)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}

func base64Decode(s string) ([]byte, error) {
	// add back missing padding
	switch len(s) % 4 {
	case 2:
		s += "=="
	case 3:
		s += "="
	}
	return base64.URLEncoding.DecodeString(s)
}
