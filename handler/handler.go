package handler

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// 返回上传html页面
		data, err := ioutil.ReadFile("./static/view/upload1.html")
		if err != nil {
			io.WriteString(w, "internel server error")
			return
		}
		io.WriteString(w, string(data))
		// 另一种返回方式:
		// 动态文件使用http.HandleFunc设置，静态文件使用到http.FileServer设置(见main.go)
		// 所以直接redirect到http.FileServer所配置的url
		// http.Redirect(w, r, "/static/view/index.html",  http.StatusFound)
	} else if r.Method == "POST" {
		// 接收文件流及存储到本地目录
		file, head, err := r.FormFile("file")
		if err != nil {
			fmt.Printf("Failed to get data, err: %s", err.Error())
			return
		}
		defer file.Close()
		newFile, err := os.Create("D:/FileStore/data/" + head.Filename)
		if err != nil {
			fmt.Printf("Failed to create file, err: %s", err.Error())
			return
		}
		defer newFile.Close()

		_, err = io.Copy(newFile, file)
		if err != nil {
			fmt.Printf("Failed to save file, err: %s", err.Error())
			return
		}
		http.Redirect(w, r, "/file/upload/suc", http.StatusFound)

	}
}

func UploadSucHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Upload Finished!")

}
