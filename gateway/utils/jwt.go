/*
 * @Author       : jayj
 * @Date         : 2021-06-24 09:40:59
 * @Description  :
 */
package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
)

const (
	encryptSalt = "tomatobell" // salt
	expire_time = 3 * time.Hour
)

var jwtSecret = []byte(encryptSalt)

type Claims struct {
	Userid   uint   `json:"userid"`
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// GenerateToken 生成token
func GenerateToken(userid uint, username, password string) (token string, err error) {

	now := time.Now()
	expireTime := now.Add(expire_time)

	claims := Claims{
		userid,
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-blog",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = tokenClaims.SignedString(jwtSecret)

	return token, err
}

// ParseToken 通过token获得claims
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		} else if ok {
			return claims, err
		}
	}

	return nil, err
}
