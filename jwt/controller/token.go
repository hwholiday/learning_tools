package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type TokenController struct {
	BaseController
}

func (c *TokenController) CreateToken(w http.ResponseWriter, r *http.Request) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims, 4)
	claims["exp"] = time.Now().Add(30 * time.Second).Unix()
	claims["uid"] = 123
	claims["name"] = "howie"
	claims["iat"] = time.Now().Unix()
	token.Claims = claims
	t, err := token.SignedString([]byte(c.GetJwtKey()))
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(t))
	return
}

func (c *TokenController) TestToken(w http.ResponseWriter, r *http.Request) {
	authString := r.Header.Get("Authorization")
	token, err := jwt.Parse(authString, func(token *jwt.Token) (interface{}, error) {
		return []byte(c.GetJwtKey()), nil
	})
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	if !token.Valid {
		w.Write([]byte("token 不合法"))
		return
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		fmt.Println(claims["name"])
		fmt.Println(claims["uid"])
	}
	w.Write([]byte("token 合法"))
	return
}

func (c *TokenController) CreateTokenByRsa(w http.ResponseWriter, r *http.Request) {
	token := jwt.New(jwt.SigningMethodRS256)
	claims := make(jwt.MapClaims, 4)
	claims["exp"] = time.Now().Add(30 * time.Second).Unix()
	claims["uid"] = 123
	claims["name"] = "howie"
	claims["iat"] = time.Now().Unix()
	token.Claims = claims
	signBytes, err := ioutil.ReadFile("/home/howie/go/src/test/jwt/conf/rsa_private_key.pem")
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	t, err := token.SignedString(signKey)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(t))
	return

}

func (c *TokenController) TestRsaToken(w http.ResponseWriter, r *http.Request) {
	authString := r.Header.Get("Authorization")
	token, err := jwt.Parse(authString, func(token *jwt.Token) (interface{}, error) {
		if token.Method == jwt.SigningMethodRS256 {
			fmt.Println("使用了相同的加密")
		} else {
			fmt.Println("没有使用相同的加密")
		}
		signBytes, err := ioutil.ReadFile("/home/howie/go/src/test/jwt/conf/rsa_public_key.pem")
		if err != nil {
			return nil, err
		}
		return jwt.ParseRSAPublicKeyFromPEM(signBytes)
	})
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	if !token.Valid {
		w.Write([]byte("token 不合法"))
		return
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		fmt.Println(claims["name"])
		fmt.Println(claims["uid"])
	}
	w.Write([]byte("token 合法"))
	return
}
