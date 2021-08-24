package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/nahidhasan98/jwt_practice/db"
	"github.com/nahidhasan98/jwt_practice/model"
)

var secretKey = "somethingVerySecret"

func generateRefreshToken(username string) (string, error) {
	claims := jwt.MapClaims{
		"username":  username,
		"exp":       time.Now().Add(time.Minute * (24 * 60)).Unix(),
		"tokenType": "refresh",
	}

	tkn := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tkn.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return token, nil
}

func generateAccessToken(username string) (string, error) {
	claims := jwt.MapClaims{
		"username":  username,
		"exp":       time.Now().Add(time.Minute * 15).Unix(),
		"tokenType": "access",
	}

	tkn := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tkn.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return token, nil
}

func prepareTokens(user *model.User) (*model.Token, error) {
	accessToken, err := generateAccessToken(user.Username)
	if err != nil {
		return nil, err
	}

	refreshToken, err := generateRefreshToken(user.Username)
	if err != nil {
		return nil, err
	}

	tkn := &model.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return tkn, nil
}

func Login(w http.ResponseWriter, r *http.Request) {
	var user *model.User

	//receiving request body
	rBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(rBody, &user)
	if err != nil {
		//fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if user.Username != db.User.Username || user.Password != db.User.Password {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}

	var token *model.Token

	token, err = prepareTokens(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//response back to client
	w.Header().Set("content-type", "application/json")
	data, err := json.Marshal(token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(data)
}
