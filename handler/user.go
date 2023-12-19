package handler

import (
	"FileStore/db"
	"FileStore/util"
	"io/ioutil"
	"net/http"
)

const pwd_salt = "*#890"

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
		w.Write([]byte("Success"))
	} else {
		w.Write([]byte("Failed"))
	}

}
