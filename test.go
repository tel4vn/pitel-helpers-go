package main

import (
	"fmt"
)

var TEST_API_KEY string = "81EdYnwrvimH"
var TEST_SECRET_KEY string = "Ib19uY8v"
var TEST_USERID string = "1001"

func main() {
	setKeys(TEST_API_KEY, TEST_SECRET_KEY, TEST_USERID)
	token := getAccessToken()

	fmt.Println(token)
}
