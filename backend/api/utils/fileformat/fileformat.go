package fileformat

import (
	"path"
	"strings"

	"github.com/twinj/uuid"
)

//Function making for generate fo unique file name for upload
func UniqueFormat(fn string) string {
	//path.Ext() get the extension of the file
	fileName := strings.TrimSuffix(fn, path.Ext(fn))
	extension := path.Ext(fn)
	u := uuid.NewV4()
	newFileName := fileName + "-" + u.String() + extension

	return newFileName

}
