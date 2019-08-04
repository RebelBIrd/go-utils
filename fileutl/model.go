package fileutl

import (
	"fmt"
	"github.com/qinyuanmao/go-utils/strutl"
	"io"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

type FileType int

const (
	_ FileType = iota
	FT_VIDEO
	FT_IMAGE
	FT_DOC
	FT_APP
	FT_ZIP
	FT_MUSIC
	FT_OTHRER
)

type FileInfo struct {
	Name     string
	Suffix   string
	Type     FileType
	Category string
	Path     string
	Md5      string
}

func GinFileHandler(v *multipart.FileHeader, path, fileName string) (info *FileInfo) {
	if fileName == "" {
		fileName = v.Filename
	}
	if fileName == "" {
		return
	} else {
		rename, fileSuffix := DecodeFileName(fileName)
		fileType, category := GetFileType(fileName)
		info = &FileInfo{
			Name:     rename,
			Path:     path + "/" + category + "/" + fileName,
			Suffix:   fileSuffix,
			Category: category,
			Type:     fileType,
		}
		PathExistOrCreate(path)
		PathExistOrCreate(path + "/" + category)
		if PathExist(info.Path) {
			info.Name = rename + strutl.GetRandomString(4)
			info.Path = path + "/" + category + "/" + info.Name + fileSuffix
		}
		file, _ := v.Open()
		defer file.Close()
		out, _ := os.Create(info.Path)
		defer out.Close()
		if _, err := io.Copy(out, file); err != nil {
			fmt.Println(err.Error())
			return nil
		}
		return
	}
}

func PathExistOrCreate(path string) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		_ = os.Mkdir(path, os.ModePerm)
	}
}

func PathExist(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

func DecodeFileName(filePath string) (fileName, fileType string) {
	filePath = path.Base(filePath)
	fileType = path.Ext(filePath)
	fileName = strings.TrimSuffix(filePath, fileType)
	switch fileType {
	case ".gz":
		if strings.HasSuffix(fileName, ".tar") {
			fileType = ".tar.gz"
			fileName = strings.TrimSuffix(fileName, ".tar")
		}
	case ".exe":
		if strings.HasSuffix(fileName, ".asp") {
			fileType = ".asp.exe"
			fileName = strings.TrimSuffix(fileName, ".exe")
		}
	}
	return
}

func GetFileType(fileName string) (FileType, string) {
	_, fileType := DecodeFileName(fileName)
	fileType = strings.ToLower(fileType)[1:]
	switch fileType {
	case "avi", "mov", "rmvb", "fmv", "m4", "3gp", "mkv", "f4v":
		return FT_VIDEO, "video"
	case "bmp", "jpg", "png", "ico", "tif", "gif", "pcx", "tga", "exif", "fpx", "svg", "psd", "cdr", "pcd", "dxf", "ufo", "eps", "ai", "raw", "WMF", "webp", "jpeg":
		return FT_IMAGE, "image"
	case "txt", "doc", "docx", "ppt", "pptx", "wpf", "md", "xls", "xlsx", "pdf":
		return FT_DOC, "doc"
	case "apk", "app", "exe", "ipa":
		return FT_APP, "app"
	case "zip", "tar.gz", "dmg", "rar", "tar", "7z", "iso", "bz2":
		return FT_ZIP, "zip"
	case "cd", "aiff", "mp3", "wma", "ogg", "acc", "amr", "mid":
		return FT_MUSIC, "music"
	default:
		return FT_OTHRER, "other"
	}
}
