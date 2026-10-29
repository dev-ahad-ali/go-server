package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
)

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type Payload struct {
	Sub         int    `json:"sub"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	IsShopOwner string `json:"is_shop_owner"`
}

func CreateJwt(secret string, data Payload) (string, error) {
	header := Header{
		Alg: "HS256",
		Typ: "JWT",
	}

	headerByteArr, err := json.Marshal(header)
	if err != nil {
		return "", err
	}

	headerB64 := base64UrlEncoding(headerByteArr)

	dataByteArr, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	payloadB64 := base64UrlEncoding(dataByteArr)

	message := headerB64 + "." + payloadB64

	secretByteArr := []byte(secret)
	messageByteArr := []byte(message)

	h := hmac.New(sha256.New, secretByteArr)
	h.Write(messageByteArr)

	signature := h.Sum(nil)
	signatureB64 := base64UrlEncoding(signature)

	jwt := headerB64 + "." + payloadB64 + "." + signatureB64

	return jwt, nil
}

func base64UrlEncoding(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}
