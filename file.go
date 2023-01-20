package main

import (
	"github.com/h2non/filetype"
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
