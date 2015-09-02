package main

import (
	"net/http"
	"os"
	"crypto/md5"
	"io"
	"time"
	"fmt"
)


func getLibs(rw http.ResponseWriter, req *http.Request) {
	log.Info("getFile")
	http.ServeFile(rw, req, dataDir+"libs")
}

func createLibs(rw http.ResponseWriter, req *http.Request) {
	log.Info("createFile")
	fileName := "libs"

	// make the directory
	os.MkdirAll((dataDir), 0777)

	file, err := os.Create(dataDir + fileName)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(err.Error()))
		return
	}
	defer file.Close()

	hash := md5.New()
	multiWriter := io.MultiWriter(hash, file)

	size, err := io.Copy(multiWriter, req.Body)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(err.Error()))
		return
	}
	writeBody(object{Name: fileName, ModTime: time.Now(), CheckSum: fmt.Sprintf("%x", hash.Sum(nil)), Size: size}, rw, 200)
}

