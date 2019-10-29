package fileutl

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/qinyuanmao/go-utils/logutl"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

type Manager struct {
	Name string
	Path string
}

func (this Manager) Write(content string) {
	buffer := []byte(content)
	this.Open(func(e error, file *os.File) {
		if e == nil {
			_, _ = file.Write(buffer)
		}
	})
}

func (this Manager) WriteBase64(content string) {
	if !this.IsExist() {
		_ = this.Create()
	}
	buffer, _ := base64.StdEncoding.DecodeString(content)
	this.Open(func(e error, file *os.File) {
		if e == nil {
			_, _ = file.Write(buffer)
		}
	})
}

func (this *Manager) Create() error {
	if !strings.HasSuffix(this.Path, "/") {
		this.Path += "/"
	}
	file, err := os.Create(this.Path + this.Name)
	if err != nil {
		println("create file error", err)
	}
	defer file.Close()
	return err
}

func (this *Manager) Open(block func(error, *os.File)) {
	file, err := os.OpenFile(this.Path+this.Name, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err.Error())
		block(err, nil)
	} else {
		block(nil, file)
	}
	defer file.Close()
}

func (file Manager) IsExist() bool {
	_, err := os.Stat(file.Path)
	if err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	} else {
		return false
	}
}

func (file Manager) PathExistOrCreate() {
	_, err := os.Stat(file.Path)
	if os.IsNotExist(err) {
		_ = os.Mkdir(file.Path, os.ModePerm)
	}
}

func (this Manager) ReadAll(block func([]byte, error)) {
	this.Open(func(e error, file *os.File) {
		if e == nil {
			block(ioutil.ReadAll(file))
		}
	})
}
func (this *Manager) DeleteFile() error {
	return os.Remove(this.Path)
}

func DeleteFile(path string) error {
	return os.Remove(path)
}

func (this Manager) GetMd5() string {
	file, err := os.Open(this.Path)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	defer file.Close()
	md5h := md5.New()
	_, err = io.Copy(md5h, file)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	} else {
		return hex.EncodeToString(md5h.Sum(nil))
	}
}

func OpenFile(filePath string) error {
	switch runtime.GOOS {
	case "darwin":
		return exec.Command("open", filePath).Start()
	case "windows":
		cmd := exec.Command("cmd", "/k", "start", filePath)
		//cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		return cmd.Start()
	case "linux":
		return exec.Command("xdg-open", filePath).Start()
	default:
		return errors.New("OS is not defined!")
	}
}

func OpenPath(path string, fileName string) error {
	if fileName == "" {
		switch runtime.GOOS {
		case "darwin":
			return exec.Command("open", path).Start()
		case "windows":
			cmd := exec.Command("cmd", "/k", "explorer", path)
			// cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
			return cmd.Start()
		case "linux":
			return exec.Command("nautilus", path).Start()
		default:
			return nil
		}
	} else {
		switch runtime.GOOS {
		case "darwin":
			return exec.Command("open", "-R", path+fileName).Start()
		case "windows":
			cmd := exec.Command("cmd", "/k", "explorer.exe", "/select,"+path+fileName)
			// cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
			return cmd.Start()
		case "linux":
			return exec.Command("nautilus", path+fileName).Start()
		default:
			return nil
		}
	}
}

func OpenFilePath(filePath string) error {
	switch runtime.GOOS {
	case "darwin":
		return exec.Command("open", "-R", filePath).Start()
	case "windows":
		cmd := exec.Command("cmd", "/k", "explorer.exe", "/select,"+filePath)
		// cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		return cmd.Start()
	case "linux":
		return exec.Command("nautilus", filePath).Start()
	default:
		return nil
	}
}

func ReadFile(filePath string) string {
	r, _ := ReadFileLines(filePath)
	return strings.Join(r, "\n")
}

func ReadFileLines(filePath string) ([]string, error) {
	result := []string{}
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		logutl.Error(err)
		return result, err
	}
	s := string(b)
	for _, lineStr := range strings.Split(s, "\n") {
		lineStr = strings.TrimSpace(lineStr)
		if lineStr == "" {
			continue
		}
		result = append(result, lineStr)
	}
	return result, nil

}
