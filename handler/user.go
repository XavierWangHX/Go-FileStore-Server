package handler

import (
	"FileStore/db"
	"FileStore/util"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const pwd_salt = "*#890"

func SignInHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		data, err := ioutil.ReadFile("./static/view/signin.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(data)
		return
	}
	r.ParseForm()
	username := r.Form.Get("username")
	passwd := r.Form.Get("password")
	enc_passwd := util.Sha1([]byte(passwd + pwd_salt))
	passwd_check := db.UserSignin(username, enc_passwd)
	if !passwd_check {
		w.Write([]byte("SignIn Failed"))
		return
	}
	token := GenToken(username)
	res := db.UpdateToken(username, token)
	if !res {
		w.Write([]byte("Update Token Failed"))
		return
	}
	resp := util.RespMsg{
		Code: 0,
		Msg:  "OK",
		Data: struct {
			Location string
			Username string
			Token    string
		}{
			Location: "http://" + r.Host + "/static/view/home.html",
			Username: username,
			Token:    token,
		},
	}
	w.Write(resp.JSONBytes())
}

func UserInfoHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.Form.Get("username")
	/*token := r.Form.Get("token")
	istokenvalid := IsTokenValid(token)
	if !istokenvalid {
		w.WriteHeader(http.StatusForbidden)
		return
	}*/
	user, err := db.GetUserInfo(username)
	if err != nil {
		fmt.Println("Failed to get user info")
		w.WriteHeader(http.StatusForbidden)
		return
	}
	resp := util.RespMsg{
		Code: 0,
		Msg:  "OK",
		Data: user,
	}
	w.Write(resp.JSONBytes())

}

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		data, err := ioutil.ReadFile("./static/view/signup.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(data)
		return
	}
	r.ParseForm()
	username := r.Form.Get("username")
	passwd := r.Form.Get("password")
	if len(username) < 3 || len(passwd) < 5 {
		w.Write([]byte("Invalid Parameter"))
		return
	}
	enc_pwd := util.Sha1([]byte(passwd + pwd_salt))
	suc := db.UserSignUp(username, enc_pwd)
	if suc {
		resp := util.RespMsg{
			Code: 0,
			Msg:  "OK",
			Data: "http://" + r.Host + "/user/signin",
		}

		w.Write(resp.JSONBytes())
	} else {
		w.Write([]byte("Failed"))
	}

}
func GenToken(username string) string {
	// 40位字符:md5(username+timestamp+token_salt)+timestamp[:8]
	ts := fmt.Sprintf("%x", time.Now().Unix())
	tokenPrefix := util.MD5([]byte(username + ts + "_tokensalt"))
	return tokenPrefix + ts[:8]
}

func IsTokenValid(token string) bool {
	if len(token) != 40 {
		return false
	}
	return true
}
