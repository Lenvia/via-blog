package model

import (
	"context"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"mime/multipart"
	"via-blog/utils"
	"via-blog/utils/errmsg"
)

var Zone = utils.Zone
var AccessKey = utils.AccessKey
var SecretKey = utils.SecretKey
var Bucket = utils.Bucket
var ImgUrl = utils.QiniuServer

// UploadFile 上传文件函数
func UploadFile(file multipart.File, fileSize int64) (string, int) {
	putPolicy := storage.PutPolicy{
		Scope: Bucket,
	}
	mac := qbox.NewMac(AccessKey, SecretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := setConfig()

	putExtra := storage.PutExtra{}

	formUploader := storage.NewFormUploader(&cfg)

	ret := storage.PutRet{}

	err := formUploader.PutWithoutKey(context.Background(), &ret, upToken, file, fileSize, &putExtra)
	if err != nil {
		return "", errmsg.ERROR
	}
	url := ImgUrl + ret.Key

	println(ImgUrl)
	return url, errmsg.SUCCESS
}


func setConfig() storage.Config {
	config := storage.Config{
		Zone:          selectZone(Zone),
		UseCdnDomains: false,
		UseHTTPS:      false,
	}
	return config
}

func selectZone(id int) *storage.Zone {
	switch id {
	case 1:
		return &storage.ZoneHuadong
	case 2:
		return &storage.ZoneHuabei
	case 3:
		return &storage.ZoneHuanan
	default:
		return &storage.ZoneHuadong
	}
}
