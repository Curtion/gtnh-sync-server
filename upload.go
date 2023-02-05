package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func upload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var req Request
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			log.Println(err)
		}
		token := req.Token
		if verifyToken(token) {
			log.Println(token, "验证成功")
		} else {
			log.Println(token, "校验失败")
			result := ResultErr{Code: 403, Msg: "校验失败", Data: "error"}
			msg, _ := json.Marshal(result)
			w.Write(msg)
		}
	} else {
		result := ResultErr{Code: 400, Msg: "当前仅支持POST方法请求", Data: "error"}
		msg, _ := json.Marshal(result)
		w.Write(msg)
	}
}
