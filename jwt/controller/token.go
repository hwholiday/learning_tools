package controller

import (
	"net/http"
	"time"
	"fmt"
	"github.com/dgrijalva/jwt-go"
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
	if claims,ok:=token.Claims.(jwt.MapClaims);ok{
		fmt.Println(claims["name"])
		fmt.Println(claims["uid"])
	}
	w.Write([]byte("token 合法"))
	return
}
