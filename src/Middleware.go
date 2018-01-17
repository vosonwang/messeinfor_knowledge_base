package main

import (
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"fmt"
	"time"
	"github.com/go-redis/redis"
	"messeinfor.com/messeinfor_knowledge_base/src/model"
	"messeinfor.com/messeinfor_knowledge_base/src/conf"
	"messeinfor.com/messeinfor_knowledge_base/src/handler"
	"github.com/satori/go.uuid"
	"encoding/json"
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
				/*如果在redis上未过期，则续命2小时*/
				err = client.Set(username, "", 2*time.Hour).Err()
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
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "无法解析用户信息！")
	}
	if p := model.FindUser(user); p != nil {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"exp": time.Now().Add(time.Hour * time.Duration(8)).Unix(), //token在8小时后过期
			"iat": time.Now().Unix(),
			"sub": user.Username,
		})

		tokenString, err := token.SignedString([]byte(conf.SecretKey))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, "Error while signing the token")
		}


		response := Token{tokenString, (*p).Username, (*p).Id}
		/*在redis上设置token为2小时过期*/
		err = client.Set(user.Username, "", 2*time.Hour).Err()
		if err != nil {
			panic(err)
		}
		handler.JsonResponse(w, response)

	} else {
		handler.JsonResponse(w, "用户名密码错误！")
	}

}
