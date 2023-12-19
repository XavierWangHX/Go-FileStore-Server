package main

import (
	"FileStore/handler"
	"fmt"
)
import "net/http"

func main() {
	http.HandleFunc("/file/upload", handler.UploadHandler)
	http.HandleFunc("/file/upload/suc", handler.UploadSucHandler)
	http.HandleFunc("/file/meta", handler.GetFileMetaHandler)
	http.HandleFunc("/file/download", handler.DownloadHandler)
	http.HandleFunc("/file/update", handler.FileMetaUpdateHandler)
	http.HandleFunc("/file/delete", handler.FileDeleteHandler)
	http.HandleFunc("/user/signup", handler.SignUpHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("failed to start server, err:%s", err.Error())
	}
}
