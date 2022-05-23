package controllers

import (
	"encoding/json"
	"net/http"
	"testando/api/models"
	"testando/api/utils"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var JwtKey = []byte("secret_key")

/* var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}
*/
type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

/* type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
} */

func Login(w http.ResponseWriter, r *http.Request) {
	var err error
	var credentials Credentials
	err = json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := models.GetUserByEmail(credentials.Email)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), ([]byte(credentials.Password)))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//expirationTime := time.Now().Add(time.Minute * 5)

	/* 	claims := &Claims{
		Username: credentials.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}  */

	//token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, err := token.SignedString(JwtKey)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	models.UpdateTokenByEmail(tokenString, user)

	utils.Response(w, tokenString, http.StatusOK)

	/* http.SetCookie(w,
	&http.Cookie{
		Name:  "token",
		Value: tokenString,
		//Expires: expirationTime,
	}) */
}

func Refresh(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tokenStr := cookie.Value

	//claims := &Claims{}

	/* tkn, err := jwt.ParseWithClaims(tokenStr, claims,
	func(t *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	}) */
	tkn, err := jwt.Parse(tokenStr,
		func(t *jwt.Token) (interface{}, error) {
			return JwtKey, nil
		})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }

	//expirationTime := time.Now().Add(time.Minute * 5)

	//claims.ExpiresAt = expirationTime.Unix()

	/* token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) */
	/* 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	 */token := jwt.New(jwt.SigningMethodHS256)
	tokenString, err := token.SignedString(JwtKey)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w,
		&http.Cookie{
			Name:  "refresh_token",
			Value: tokenString,
			///*  */Expires: expirationTime,
		})
}
