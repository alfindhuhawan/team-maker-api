package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"strconv"
	"team-maker-api/shared/model/repository"
	"time"

	"github.com/golang-jwt/jwt"
)

func Auth(keyConfig string, jwtSecret string, hexID string, userSession *repository.UserSession) (string, error) {
	var tokenAuth string

	jsonString, err := json.Marshal(userSession) // Set Data
	if err != nil {
		return "", errors.New("Error encoding JSON")
	}

	key := []byte(keyConfig)

	// ENCRYPT
	c, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}
	encryptByte := gcm.Seal(nonce, nonce, jsonString, nil)

	// encryptByte, err = helper.Encrypt(jsonString, key)
	// if err != nil {
	// 	return nil, err
	// }

	pxl := hex.EncodeToString(encryptByte)

	currentTimestamp := time.Now().UTC().Unix()
	var ttl int64 = (3600 * 10) // expired time in second
	// md5 of sub & iat
	h := md5.New()
	io.WriteString(h, hexID)
	io.WriteString(h, strconv.FormatInt(int64(currentTimestamp), 10))
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.

	subRandom, _ := GenerateRandomString(`[A-Z0-9]{7}-[A-Z0-9]{7}-[A-Z0-9]{7}-[A-Z0-9]{7}`)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": subRandom,
		"iat": currentTimestamp,
		"exp": currentTimestamp + ttl, // 0 <<<< "0" for unlimited expired
		"nbf": currentTimestamp,
		"jti": h.Sum(nil),
		"pxl": pxl,
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenAuth, err = token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return tokenAuth, nil
}
