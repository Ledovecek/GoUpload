package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	router := gin.Default()
	router.Static("/download", "download")
	router.POST("/", func(c *gin.Context) {
		if c.Request.ContentLength <= AllowedFileSize {
			file, err := c.FormFile("file")
			err = c.SaveUploadedFile(file, "download/"+file.Filename)
			readFile, err := os.ReadFile("download/" + file.Filename)
			if err != nil {
				panic(err)
				return
			}
			uploadedFile := New(readFile, file.Size)
			fileExtension := filepath.Ext("download/" + file.Filename)
			if IsAllowed(uploadedFile) {
				generatedUuid := uuid.NewV4().String()
				os.Rename("download/"+file.Filename, "download/"+generatedUuid+filepath.Ext("download/"+fileExtension))
				c.String(http.StatusOK, fmt.Sprintf("Link: http://localhost:8080/download/%s"+fileExtension, generatedUuid))
			} else {
				c.String(http.StatusNotAcceptable, "This file format is not allowed to upload")
			}
		} else {
			c.Abort()
			c.String(http.StatusNotAcceptable, "File is to large")
		}
	})
	err := router.Run()
	if err != nil {
		panic("Something went really wrong")
		return
	}
}
