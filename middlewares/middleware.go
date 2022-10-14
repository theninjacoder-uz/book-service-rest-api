package middlewares

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"task/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetMiddlewareAuthentication(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		//Get auth key and sign from request header
		key := c.Request.Header.Get("Key")
		sign := c.Request.Header.Get("Sign")

		//Get user's secret string from database
		userSecret, err := getUserSecret(db, key)
		if err != nil {
			models.ErrorResponse(c, http.StatusUnauthorized, errors.New("Unauthorized"))
			c.Abort()
			return
		}

		// Build sign string from request uri and body and generate signature
		proto := "http://"
		if strings.HasPrefix(c.Request.Proto, "HTTPS") {
			proto = "https://"
		}
		uri := fmt.Sprintf("%v%v%v%v", c.Request.Method, proto, c.Request.Host, c.Request.URL)
		fmt.Println(uri)

		// Get request body
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			models.ErrorResponse(c, http.StatusUnauthorized, errors.New("Error when read request body"))
			c.Abort()
			return
		}

		// Generate valid sign
		hash, err := generateSign(body, uri, userSecret)
		if err != nil {
			models.ErrorResponse(c, http.StatusUnauthorized, errors.New("Error when generating sign"))
			c.Abort()
			return
		}

		// Check if client's sign is valid or not
		fmt.Printf("hash: %s , sign: %s\n", hash, sign)
		if hash != sign {
			models.ErrorResponse(c, http.StatusUnauthorized, errors.New("Unauthorized"))
			c.Abort()
			return
		}

		//Reassign the body again
		c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
		c.Next()
	}
}

func getUserSecret(db *gorm.DB, key string) (string, error) {
	user := models.User{}
	savedUser, err := user.GetUserInfo(db, key)
	if err != nil {
		return "", errors.New("invalid key")
	}
	return savedUser.Secret, nil
}

func generateSign(body []byte, uri string, secret string) (string, error) {
	compactedBuffer := new(bytes.Buffer)
	if err := json.Compact(compactedBuffer, body); err != nil {
		return "", errors.New("Error during body parse")
	}
	rawSign := uri + compactedBuffer.String() + secret
	hash := encode(rawSign)
	return hash, nil
}

func encode(data string) string {
	h := md5.New()
	io.WriteString(h, data)
	return fmt.Sprintf("%x", h.Sum(nil))
}
