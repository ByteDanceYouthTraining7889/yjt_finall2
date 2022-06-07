package service

import (
	"douyin/pkg"
	"errors"
	"mime/multipart"
	"os"
)

type FileInfo struct {
	Name      string
	AccessUrl string
}

func UploadFile(fileType pkg.FlieType, file multipart.File, fileHeader *multipart.FileHeader) (*FileInfo, error) {
	fileName := pkg.GetFileName(fileHeader.Filename)
	uploadSavePath := pkg.GetSavePath()
	dst := uploadSavePath + "/" + fileName
	if !pkg.CheckContainExit(fileType, fileName) {
		return nil, errors.New("file suffix is not supported. ")
	}
	if pkg.CheckSavePath(uploadSavePath) {
		if err := pkg.CreateSavePath(uploadSavePath, os.ModePerm); err != nil {
			return nil, errors.New("failed to create save directory. ")
		}
	}
	// if pkg.CheckMaxSize(fileType, file) {
	// 	return nil, errors.New("exceeded maximum file limit. ")
	// }
	if pkg.CheckPermission(uploadSavePath) {
		return nil, errors.New("insufficient file permissions. ")
	}
	if err := pkg.SaveFile(fileHeader, dst); err != nil {
		return nil, err
	}
	accessUrl := "http://10.17.119.65/static" + "/" + fileName
	return &FileInfo{
		Name:      fileName,
		AccessUrl: accessUrl,
	}, nil
}
