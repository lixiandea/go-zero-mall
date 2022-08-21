/**
* @Author:Dijiang
* @Description:
* @Date: Created in 21:46 2022/6/17
* @Modified By: Dijiang
 */

package jwtx

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v4"
)

func GetToken(secretKey string, iat, seconds, uid int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["uid"] = uid
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

func GetClaims(secretKey string, tokenString string) (claims *jwt.Claims, err error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(secretKey), nil
	}, jwt.WithJSONNumber())
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("token not valid! ")
	}
	return &token.Claims, err
}
func GetUserId(secretKey string, tokenString string) (uid int64, err error) {
	claimsPtr, err := GetClaims(secretKey, tokenString)
	if err != nil {
		return -1, err
	}
	mapClaims, ok := (*claimsPtr).(jwt.MapClaims)
	if !ok {
		return -1, errors.New("not a map claims")
	}
	// thr type of mapClaims["uid"] is float64 while it cames from type of int64
	uid, _ = mapClaims["uid"].(json.Number).Int64()
	// uid = int64(math.Round(mapClaims["uid"].(float64)))
	return uid, err
}
