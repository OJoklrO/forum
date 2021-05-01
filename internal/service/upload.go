package service

import (
	"errors"
	"forum/global"
	"github.com/google/uuid"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"path/filepath"
	"strconv"
)

func (svc *Service) Upload(file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	ext := path.Ext(fileHeader.Filename)
	// check file type
	allow := false
	for _, v := range global.AppSetting.UploadExtensions {
		if v == ext {
			allow = true
			break
		}
	}
	if !allow {
		return "", errors.New("the file type is not allowed")
	}

	// unique file name
	var dstPath, fileName string
	for {
		fileName = uuid.New().String() + ext
		dstPath = filepath.Join(global.AppSetting.UploadSavePath, fileName)
		_, err := os.Stat(dstPath)
		if os.IsNotExist(err) {
			break
		}
	}

	err := os.MkdirAll(global.AppSetting.UploadSavePath, os.ModePerm)
	if err != nil {
		return "", err
	}

	// size
	content, _ := ioutil.ReadAll(file)
	if len(content) > global.AppSetting.UploadMaxSize*1024*1024 {
		return "", errors.New("exceeded max file size: " +
			strconv.Itoa(global.AppSetting.UploadMaxSize) + "M")
	}

	err = ioutil.WriteFile(dstPath, content, os.ModePerm)
	if err != nil {
		return "", err
	}

	return global.ServerSetting.Url + global.AppSetting.UploadApi + "/" +
		fileName, nil
}
