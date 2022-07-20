/**
* @Author:Dijiang
* @Description:
* @Date: Created in 21:46 2022/6/17
* @Modified By: Dijiang
 */

package jwtx

import "github.com/form3tech-oss/jwt-go"

func GetToken(secretKey string, iat, seconds, uid int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["uid"] = uid
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
