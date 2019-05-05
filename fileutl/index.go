package fileutl

import (
	"fmt"
	"io/ioutil"
	"os"
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

func (this *Manager) Create() error {
	file, err := os.Create(this.Path + this.Name)
	if err != nil {
		println("create file error", err)
	}
	defer file.Close()
	return err
}

func (this *Manager) Open(block func(error, *os.File)) {
	file, err := os.OpenFile(this.Path + this.Name, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
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