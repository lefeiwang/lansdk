package tool

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

func ReadFile(srcFile string) (content []byte, err error) {
	if _, err := os.Stat(srcFile); err != nil {
		return nil, err
	}
	content, err = ioutil.ReadFile(srcFile)
	if err != nil {
		return nil, err
	}

	return content, nil
}

func WriteFile(srcFile string, content []byte) (err error) {
	err = ioutil.WriteFile(srcFile, content, 0666)
	if err != nil {
		return err
	}
	return nil
}

// 判断所给路径文件/文件夹是否存在
func FileIsExist(srcFile string) bool {
	_, err := os.Stat(srcFile)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

// 删除所给路径文件/文件夹
func RemoveFile(srcFile string) error {
	err := os.RemoveAll(srcFile)
	return err
}

// 创建嵌套文件夹
func MkDir(srcFile string) error {
	err := os.MkdirAll(srcFile, os.ModePerm)
	return err
}

func CopyFile(srcFile, destFile string) (int, error) {
	fileDir, _ := filepath.Split(destFile)
	err := MkDir(fileDir)
	if err != nil {
		return 0, err
	}
	input, err := ReadFile(srcFile)
	if err != nil {
		return 0, err
	}

	err = WriteFile(destFile, input)
	if err != nil {
		return 0, err
	}

	return len(input), nil
}

func CopyFileBuff(src, dst string, BUFFERSIZE int64) (cn int, err error) {
	fileDir, _ := filepath.Split(dst)
	err = MkDir(fileDir)
	if err != nil {
		return 0, err
	}
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return
	}
	if !sourceFileStat.Mode().IsRegular() {
		err = fmt.Errorf("%s is not a regular file", src)
		return
	}
	source, err := os.Open(src)
	if err != nil {
		return
	}
	defer source.Close()

	_, err = os.Stat(dst)
	if err == nil {
		err = fmt.Errorf("file %s already exists", dst)
		return
	}

	destination, err := os.Create(dst)
	if err != nil {
		return
	}
	defer destination.Close()

	if err != nil {
		panic(err)
	}

	buf := make([]byte, BUFFERSIZE)
	var rn, wn int
	for {
		rn, err = source.Read(buf)
		if err != nil && err != io.EOF {
			return
		}
		if rn == 0 {
			break
		}

		if wn, err = destination.Write(buf[:rn]); err != nil {
			return
		}
		cn += wn
	}
	return cn, nil
}
