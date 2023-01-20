package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func main() {
	router := gin.Default()
	dirEntry, _ := os.ReadDir("download")
	if dirEntry == nil {
		os.Mkdir("download", 0755)
	}
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
			if IsAllowed(uploadedFile) {
				c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
			} else {
				err := os.Remove("download/" + file.Filename)
				if err != nil {
					panic(err)
					return
				}
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
