package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

const (
	// SecretKey 这个值在生产环境中应该是私钥
	SecretKey = "general private key"
)

// Claims ...
type Claims struct {
	jwt.StandardClaims
	User  string   // 用户名
	Role  string   // 角色
	Scope []string // 权限
}

// LoginHandler ...
func LoginHandler(resp http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	byteData, _ := ioutil.ReadAll(req.Body)
	// var mapData map[string]interface{}
	mapData := map[string]interface{}{}
	err := json.Unmarshal(byteData, &mapData)
	if err != nil {
		log.Println(err)
	}
	log.Printf("认证信息: %s\n", byteData)

	username := mapData["username"].(string)
	password := mapData["password"].(string)
	if username == "admin" && password == "123456" {
		claims := Claims{
			User:  username,
			Role:  "admin",
			Scope: []string{"update", "delete"},
		}
		tokenBuilder := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		token, err := tokenBuilder.SignedString([]byte(SecretKey))
		if err != nil {
			log.Println(err)
			resp.Write([]byte(`{"msg": "token获取失败"}`))
			return
		}
		resultData := map[string]string{
			"token": token,
		}
		result, _ := json.Marshal(resultData)
		resp.Write(result)
	} else {
		resp.Write([]byte(`{"msg": "用户名/密码错误"}`))
	}
	return
}

// AuthMiddleware 中间件, 验证token合法性
func AuthMiddleware(next http.Handler) http.Handler {
	// 注意这里的类型转换
	return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		tokenString := req.Header.Get("Authorization")
		log.Println("Token: ", tokenString)
		// 解析, 验证token, 并返回这个token对象
		// jwt-go/request提供了更方便的验证请求头中token的方法, 但只针对于原生http库.
		// 考虑到我们可能会使用到beego, gin等web框架, 所以使用`jwt.ParseXXX`的方法更有用.
		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(SecretKey), nil
		})
		resultData := map[string]string{}
		if err != nil {
			resultData["msg"] = err.Error()
			result, _ := json.Marshal(resultData)
			resp.Write(result)
			return
		}

		if !token.Valid {
			resp.Write([]byte(`{"msg": "invalid token"}`))
			return
		}
		claims := token.Claims
		log.Printf("得到解析后的claims对象: %+v\n", claims)
		next.ServeHTTP(resp, req)
	})
}

// Update ...
func Update(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte(`{"msg": "success"}`))
}

/*
 * 访问localhost:7722/login, POST数据{username: admin, password: 123456}, 可以得到一个token
 * 然后访问/update接口, 在请求头中添加`Authorization: token值`, 就可以得到success
 */
func main() {
	http.HandleFunc("/login", LoginHandler)
	http.Handle("/update", AuthMiddleware(http.HandlerFunc(Update)))
	log.Println("http server started...")
	http.ListenAndServe(":7722", nil)
}
