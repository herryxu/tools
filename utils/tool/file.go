package tool

import (
	"errors"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func DirName(argv ...string) string {
	file := ""
	if len(argv) > 0 && argv[0] != "" {
		file = argv[0]
	} else {
		file, _ = exec.LookPath(os.Args[0])
	}
	path, _ := filepath.Abs(file)
	directory := filepath.Dir(path)
	return strings.Replace(directory, "\\", "/", -1)
}

func GetProPath() string {
	return DirName("root")
}

func IsExist(f string) bool {
	_, err := os.Stat(f)
	return err == nil || os.IsExist(err)
}

func IsFile(f string) bool {
	if fi, err := os.Stat(f); err != nil {
		return false
	} else {
		return !fi.IsDir()
	}
}

func IsDir(f string) bool {
	if fi, err := os.Stat(f); err != nil {
		return false
	} else {
		return fi.IsDir()
	}
}

func MkDirAll(path string) error {
	return os.MkdirAll(path, 0777)
}

func ReadFile(filename string) (string, error) {
	if !IsFile(filename) {
		return "", errors.New("file does not exist")
	}
	if ret, err := ioutil.ReadFile(filename); err == nil {
		return Bytes2String(ret), nil
	} else {
		return "", err
	}
}

func WriteFile(filename string, data string) error {
	path := DirName(filename)
	if !IsDir(path) {
		if err := MkDirAll(path); err != nil {
			return err
		}
	}
	return ioutil.WriteFile(filename, String2Bytes(data), 0777)
}

func AppendFile(filename string, data string) error {
	if !IsFile(filename) {
		return WriteFile(filename, data)
	}
	f, err := os.OpenFile(filename, os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	n, err2 := f.Seek(0, 2)
	if err2 != nil {
		return err2
	}
	_, err = f.WriteAt(String2Bytes(data), n)
	return err
}
