package main

import (
	"net/http"
	"time"
	"os"
	"fmt"
	"crypto/md5"
	"io"
	"io/ioutil"
)

func getBuild(rw http.ResponseWriter, req *http.Request) {
	log.Info("getFile")
	http.ServeFile(rw, req, dataDir+"builds/"+req.URL.Query().Get(":file"))
}

func getBuildHead(rw http.ResponseWriter, req *http.Request) {
	filename := dataDir+"builds/"+req.URL.Query().Get(":file")
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		rw.WriteHeader(http.StatusNotFound)
		return
	}
	fileStat, err := file.Stat()
	if err != nil {
		rw.WriteHeader(http.StatusNotFound)
		return
	}

	rw.Header().Set("SIZE", fmt.Sprintf("%d", fileStat.Size()))
	rw.Header().Set("NAME", fileStat.Name())
}

func createBuild(rw http.ResponseWriter, req *http.Request) {
	log.Info("createFile")
	fileName := req.URL.Query().Get(":file")

	// make the directory
	os.MkdirAll((dataDir+"builds/"), 0777)

	file, err := os.Create(dataDir+"builds/" + fileName)
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

func deleteBuild(rw http.ResponseWriter, req *http.Request) {
	log.Info("deleteFile")
	fileName := req.URL.Query().Get(":file")
	if fileName == "" {
		rw.WriteHeader(http.StatusNotFound)
		return
	}
	err := os.Remove(dataDir+"builds/"+fileName)
	if err != nil{
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(err.Error()))
		return
	}
}

func listBuilds(rw http.ResponseWriter, req *http.Request) {
	log.Info("listFiles")
	// http.ServeFile(rw, req, dataDir+"builds/"+folder)
	objects, err := builds()
	if err != nil {
		rw.WriteHeader(http.StatusNotFound)
		return
	}
	writeBody(objects, rw, http.StatusOK)
}

func builds() (objects []object, err error) {
	var filesArr []os.FileInfo
	filesArr, err = ioutil.ReadDir(dataDir+"builds/")
	if err != nil {
		return
	}
	for _, file := range filesArr {
		if !file.IsDir() {
			objects = append(objects, object{Name: file.Name(), Size: file.Size(), ModTime: file.ModTime() })
		}
	}
	return	
}
