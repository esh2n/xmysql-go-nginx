package auth

import (
	"log"
	"os"

	jwt "github.com/dgrijalva/jwt-go"
)

type JwtUser struct {
	ID   int
	Name string
	jwt.StandardClaims
}

func CreateTokenString(id int, name string) string {
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), &JwtUser{
		ID:   id,
		Name: name,
	})
	tokenstring, err := token.SignedString([]byte(os.Getenv("SIGNINGKEY")))
	if err != nil {
		log.Fatalln(err)
	}
	return tokenstring
}
