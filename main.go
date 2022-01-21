package main

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

func main() {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
		"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

}
