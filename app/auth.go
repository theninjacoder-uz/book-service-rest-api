package app

import (
	// "crypto/md5"
	"fmt"
	"io/ioutil"
	"net/http"
	// 	"task/models"
)

func ValidateSign(r *http.Request) error {
	key := r.Header.Get("Key")
	sign := r.Header.Get("Sign")

	authenticateRequest(r, key, sign)
	// return nil
	fmt.Println(key, sign)
	return nil
}

func authenticateRequest(r *http.Request, key string, secret string) (string, error) {
	uri := r.URL.RequestURI()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Println("Error when read request body")
		return "", err
	}

	fmt.Println("uri: %s , body: %s", uri, string(body))

	// sign, e := server.getUserSecret(key)
	// if e != nil {
	// 	fmt.Println("Error when get user secret")
	// 	return "", e
	// }

	// fmt.Println("Signature: %s", sign)
	// fmt.Println("Encoded data: %s", uri+string(body))
	// return uri + string(body), nil
	return "", nil
}

// func (server *Server) getUserSecret(key string) (string, error) {
// 	user := models.User{}
// 	savedUser, err := user.GetUserInfo(server.DB, key)

// 	if err != nil {
// 		return "", err
// 	}

// 	return savedUser.Secret, nil
// }

// func encode(data string) string {
// 	encodedData := md5.Sum([]byte(data))
// 	return string(encodedData[:])
// }
