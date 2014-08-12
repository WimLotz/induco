package main

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"log"
)

func marshalObjectToJson(obj interface{}) []byte {
	bytes, err := json.Marshal(obj)
	if err != nil {
		log.Printf("Json Marshalling error:%v", err)
		return nil
	}
	return bytes
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
