package entity

import "mime/multipart"

type UserImage struct {
	File multipart.File
	Handler *multipart.FileHeader
}