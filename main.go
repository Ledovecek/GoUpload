package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/h2non/filetype"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"os"
	"path/filepath"
)

const AllowedFileSize = 1092566821
const DownloadsDirectory string = "download"
const SLASH string = "/"

const (
	UnknownError  string = "An unknown error has uccured"
	SizeExceeded  string = "Maximum file size has been exceeded"
	IllegalFormat string = "Illegal file format - not archive"
	DownloadLink  string = "Link: http://localhost:8080/download/%s"
)

type uploadedFile struct {
	isArchive bool
	size      int64
}

func New(fileBytes []byte, size int64) uploadedFile {
	archive := filetype.IsArchive(fileBytes)
	return uploadedFile{isArchive: archive, size: size}
}

func IsAllowed(f uploadedFile) bool {
	if f.size <= AllowedFileSize && f.isArchive {
		return true
	} else {
		return false
	}
}

func main() {
	router := gin.Default()
	router.Static(SLASH+DownloadsDirectory, DownloadsDirectory)
	router.POST(SLASH, func(c *gin.Context) {
		if c.Request.ContentLength <= AllowedFileSize {
			file, err := c.FormFile("file")
			err = c.SaveUploadedFile(file, "download/"+file.Filename)
			readFile, err := os.ReadFile("download/" + file.Filename)
			if err != nil {
				panic(err)
				return
			}
			uploadedFile := New(readFile, file.Size)
			fileExtension := filepath.Ext(DownloadsDirectory + SLASH + file.Filename)
			if IsAllowed(uploadedFile) {
				generatedUuid := uuid.NewV4().String()
				os.Rename("download/"+file.Filename, "download/"+generatedUuid+filepath.Ext(DownloadsDirectory+SLASH+fileExtension))
				c.String(http.StatusOK, fmt.Sprintf(DownloadLink+fileExtension, generatedUuid))
			} else {
				c.String(http.StatusNotAcceptable, IllegalFormat)
			}
		} else {
			c.Abort()
			c.String(http.StatusNotAcceptable, SizeExceeded)
		}
	})
	err := router.Run()
	if err != nil {
		panic(UnknownError)
		return
	}
}
