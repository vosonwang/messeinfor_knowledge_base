package main

import (
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"fmt"
	"time"
	"github.com/go-redis/redis"
	"messeinfor.com/messeinfor_knowledge_base/src/models"
	"messeinfor.com/messeinfor_knowledge_base/src/conf"
	"messeinfor.com/messeinfor_knowledge_base/src/handler"
	"github.com/satori/go.uuid"
)

type Token struct {
	Token    string    `json:"token"`
	Username string    `json:"username"`
	UserId   uuid.UUID `json:"userId"`
}

func ValidateToken(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(conf.SecretKey), nil
		})

	if err == nil {
		claims := token.Claims.(jwt.MapClaims)
		username := claims["sub"].(string)
		_, err := client.Get(username).Result()
		/*如果token合法，并且没有过期*/
		if token.Valid {
			if err == redis.Nil {
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprint(w, "登录超时！")
			} else if err != nil {
				panic(err)
			} else {
				/*如果在redis上未过期，则续命1小时*/
				err = client.Set(username, "", time.Hour).Err()
				if err != nil {
					panic(err)
				}
				next(w, r)
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Something wrong")
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Unauthorized access to this resource")
	}
}

func NewToken(w http.ResponseWriter, r *http.Request) {
	user := models.ParseUser(r.Body)
	if a, b := models.FindUser(user); b != false {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"exp": time.Now().Add(time.Hour * time.Duration(8)).Unix(), //8小时后过期
			"iat": time.Now().Unix(),
			"sub": user.Username,
		})

		tokenString, err := token.SignedString([]byte(conf.SecretKey))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, "Error while signing the token")
		}

		response := Token{tokenString, a.Username, a.Id}
		/*如果在redis上未过期，则续命1小时*/
		err = client.Set(user.Username, "", time.Hour).Err()
		if err != nil {
			panic(err)
		}
		handler.JsonResponse(w, response)

	} else {
		w.Header().Set("Content-Type", "application/json;   charset=UTF-8")
		fmt.Fprintln(w, 0)
	}

}
