package restapi

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"team-maker-api/shared/infrastructure/util"
	"team-maker-api/shared/model/repository"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func (r *Controller) authenticated() gin.HandlerFunc {

	return func(c *gin.Context) {

		if len(c.Request.Header.Get("Authorization")) == 0 {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		tokenAuth := c.Request.Header.Get("Authorization")[7:]
		userSession, err := r.ValidateAuthToken(tokenAuth)
		if err != nil {
			c.JSON(401, gin.H{
				"code":         "401",
				"code_message": "ErrorValidateToken",
				"message":      err.Error(),
				"data":         nil,
			})
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		result := make(map[string]string, 0)
		result["username"] = userSession.Username
		result["name"] = userSession.Name
		result["role"] = userSession.Role
		result["member_code"] = userSession.MemberCode

		c.Set("user_session", result)
		c.Set("username", userSession.Username)
		c.Set("token", tokenAuth)
		c.Set("member_code", userSession.MemberCode)
		c.Set("role", userSession.Role)
		c.Set("admin_hub_code", userSession.AdminHubCode)
	}
}

func (r *Controller) authenticatedAdmin() gin.HandlerFunc {

	return func(c *gin.Context) {
		if len(c.Request.Header.Get("Authorization")) <= 7 || c.Request.Header.Get("Authorization") == "" {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		administratorOnly := "Adm1nistr4tor0nly"

		hashString := util.CreateSHA256Signature(administratorOnly)

		// GET AUTH
		tokenAuth := c.Request.Header.Get("Authorization")[7:]

		if hashString != tokenAuth {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
	}
}

// authorized is an interceptor
func (r *Controller) authorized() gin.HandlerFunc {

	return func(c *gin.Context) {

		authorized := true

		if !authorized {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
	}
}

func (r *Controller) ValidateAuthToken(reqToken string) (*repository.UserSession, error) {
	var result *repository.UserSession
	keyConfig := r.Config.AuthKey.KeyConfig
	jwtSecret := r.Config.AuthKey.JWTSecret

	token, err := jwt.Parse(reqToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(jwtSecret), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid && claims != nil {
		dataSession := claims["pxl"].(string)
		key := []byte(keyConfig)

		decoded, err := hex.DecodeString(dataSession)
		if err != nil {
			return nil, fmt.Errorf("failed to decrypt data")
		}

		plaintext, err := Decrypt(decoded, key)
		if err != nil {
			return nil, fmt.Errorf("failed to decrypt data to plaint text")
		}

		err = json.Unmarshal(plaintext, &result)
		if err != nil {
			return nil, fmt.Errorf("failed to unMarshal data")
		}
	}

	return result, nil
}

func Decrypt(ciphertext []byte, key []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, errors.New("ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	return gcm.Open(nil, nonce, ciphertext, nil)
}
