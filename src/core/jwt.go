package core

import jwt "github.com/appleboy/gin-jwt/v2"

var jwtMiddle *jwt.GinJWTMiddleware

func SetJwtMiddleware(middle *jwt.GinJWTMiddleware) {
	jwtMiddle = middle
}

func GetJwtMiddleware() *jwt.GinJWTMiddleware {
	return jwtMiddle
}
