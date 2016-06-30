package utility

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"stash.cqlcorp.net/mp/moja-portal/config"
	"stash.cqlcorp.net/mp/moja-portal/models"
	"strconv"
	"time"
)

func Debug(message string) {
	if config.Debug == "true" {
		fmt.Println(message)
	}
}
func Security(message string) {
	if config.SecurityOutput == "true" {
		fmt.Println(message)
	}
}

func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

func GenerateRandomString(s int) (string, error) {
	b, err := GenerateRandomBytes(s)
	if err != nil {
		return "", err
	}
	code := base64.URLEncoding.EncodeToString(b)
	return code[0:s], nil
}

func GetUserFromContext(c *gin.Context) (*models.User, bool) {
	// get user from context
	if userContext, ok := c.Get("user"); ok {
		if userDisplay, ok := userContext.(models.User); ok {
			return &userDisplay, true
		}
	}
	return nil, false
}

// getTimeout
func GetTimeout(timeout string) time.Duration {
	t, err := strconv.ParseInt(timeout, 10, 64)
	if err != nil {
		Debug("Cannont parse timeout. Defaulting to 15 minutes for saftey.")
		t = 15
	}
	return time.Duration(t)
}
