package upload

import (
	"qoquery/pkg/file"
	"qoquery/pkg/util"

	"fmt"
	"log"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

// 获取图片完整访问URL
func GetImageFullUrl(name string) string {
	return "127.0.0.1:8000" + "/" + GetImagePath() + name
}

// 获取图片名称
func GetImageName(name string) string {
	ext := path.Ext(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = util.EncodeMD5(fileName)
	return fileName + ext
}

// 获取图片路径
func GetImagePath() string {
	return "upload/images/"
}

// 获取图片完整路径
func GetImageFullPath() string {
	return "\\upload\\images\\"
}

// 检查图片后缀
func CheckImageExt(fileName string) bool {
	ext := file.GetExt(fileName)
	for _, allowExt := range []string{".jpg", ".jpeg", ".png"} {
		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
			return true
		}
	}
	return false
}

// 检查图片大小
func CheckImageSize(f multipart.File) bool {
	size, err := file.GetSize(f)
	if err != nil {
		log.Fatal(err)

		return false
	}
	return size <= 5
}

// 检查图片
func CheckImage(src string) error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd err: %v", err)
	}
	err = file.IsNotExistMkDir(dir + "/" + src)
	if err != nil {
		return fmt.Errorf("file.IsNotExistMkDir err: %v", err)
	}
	perm := file.CheckPermission(src)
	if perm == true {
		return fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}
	return nil
}
