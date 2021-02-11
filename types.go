package main

import (
	"encoding/json"
	"net/http"
)

type Middleware func(handlerFunc http.HandlerFunc) http.HandlerFunc

type metaData interface{}

type User struct {
	Name string `json:"name"`
	Email string`json:"email"`
	Phone int `json:"phone"`
}

func (u *User) ToJson() ([]byte, error) {
	return json.Marshal(u)
}