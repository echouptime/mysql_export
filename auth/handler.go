package auth

import (
	"encoding/base64"
	"net/http"
	"strings"
)

type AuthSecrets map[string]string

func Auth(handler http.Handler, secrets AuthSecrets) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {

		secret := request.Header.Get("Authorization")
		if !isAuth(secret, secrets) {
			//弹出输入用户名密码框
			response.Header().Set("WWW-Authenticate", `Basic realm=""`)
			response.WriteHeader(401)
			return
		}
		handler.ServeHTTP(response, request)

	})
}

func isAuth(secret string, secrets AuthSecrets) bool {
	if secrets == nil {
		return true
	}
	nodes := strings.Fields(secret)
	if len(nodes) != 2 {
		return false
	}
	plaintext, err := base64.StdEncoding.DecodeString(nodes[1])
	if err != nil {
		return false
	}
	nodes = strings.SplitN(string(plaintext), ":", 2)
	if len(nodes) != 2 {
		return false
	}

	//根据用户获取对应的密码
	password, ok := secrets[nodes[0]]
	return ok && password == nodes[1]
	//return nodes[0] == "admin" && nodes[1] == "admin123"

}
