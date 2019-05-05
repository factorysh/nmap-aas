package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

type DomainsClaims struct {
	Domains []string `json:"domains`
	jwt.StandardClaims
}

func Jwt(secret []byte, handler func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		auth, ok := r.Header["Authorization"]
		if !ok || len(auth) != 1 || len(auth[0]) < 7 || !strings.HasPrefix(strings.ToLower(auth[0]), "bearer ") {
			w.WriteHeader(401)
			return
		}
		t := auth[0][7:]
		var claims DomainsClaims
		token, err := jwt.ParseWithClaims(t, claims, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return secret, nil
		})
		if err != nil {
			w.WriteHeader(400)
			return
		}
		fmt.Println(token)
		ctx := context.WithValue(r.Context(), "claims", claims)
		handler(w, r.WithContext(ctx))
	}
}
