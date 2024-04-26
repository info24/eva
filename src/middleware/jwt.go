package middleware

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/info24/eva/model"
	"github.com/info24/eva/store"
	"log"
	"net/http"
	"time"
)

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

func NewJwt() *jwt.GinJWTMiddleware {
	middleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "GinLearn",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: "id",
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*model.User); ok {
				return jwt.MapClaims{
					jwt.IdentityKey: v.ID,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			user := &model.User{}
			user.ID = uint(claims[jwt.IdentityKey].(float64))
			return user
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var user model.User
			if err := c.ShouldBind(&user); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			tx := store.GetDB().Where("username = ? AND password = ?", user.Username, user.Password).First(&user)
			if tx.Error != nil {
				return nil, jwt.ErrFailedAuthentication
			}
			return &user, nil
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
	err = middleware.MiddlewareInit()
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
	return middleware
}
