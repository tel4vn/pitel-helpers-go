package main

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var _API_KEY string = ""
var _SECRET_KEY string = ""
var _USERID string = ""

func init() {

}
func setKeys(PITEL_API_KEY string, PITEL_SECRET_KEY string, PITEL_USERID string) {
	_API_KEY = PITEL_API_KEY
	_SECRET_KEY = PITEL_SECRET_KEY
	_USERID = PITEL_USERID

	if _API_KEY == "" {
		fmt.Errorf("PITEL_API_KEY is required")
		return
	}
	if _SECRET_KEY == "" {
		fmt.Errorf("PITEL_SECRET_KEY is required")
		return
	}

	if _USERID == "" {
		fmt.Errorf("PITEL_USERID is required")
		return
	}
}
func getAccessToken() string {
	var err error
	atClaims := jwt.MapClaims{}
	atClaims["key"] = _API_KEY
	atClaims["uid"] = _USERID
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(_SECRET_KEY))

	if err != nil {
		return ""
	}
	return token
}
