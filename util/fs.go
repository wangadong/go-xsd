package util

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

//	If a directory does not exist at the specified `dirPath`, attempts to create it.
func EnsureDirExists(dirPath string) (err error) {
	if !DirExists(dirPath) {
		if err = EnsureDirExists(filepath.Dir(dirPath)); err == nil {
			err = os.Mkdir(dirPath, os.ModePerm)
		}
	}
	return
}

//	Returns whether a directory (not a file) exists at the specified `dirpath`.
func DirExists(dirpath string) bool {
	if len(dirpath) == 0 {
		return false
	}
	stat, err := os.Stat(dirpath)
	return err == nil && stat.IsDir()
}

//	Returns whether a file (not a directory) exists at the specified `filePath`.
func FileExists(filePath string) bool {
	stat, err := os.Stat(filePath)
	return err == nil && stat.Mode().IsRegular()
}

//	Performs an `io.Copy` from the specified `io.Reader` to the specified local file.
func SaveToFile(src io.Reader, dstFilePath string) (err error) {
	var file *os.File
	if file, err = os.Create(dstFilePath); file != nil {
		defer file.Close()
		if err == nil {
			_, err = io.Copy(file, src)
		}
	}
	return
}

//	A short-hand for `ioutil.WriteFile` using `ModePerm`.
//	Also ensures the target file's directory exists.
func WriteBinaryFile(filePath string, contents []byte) error {
	EnsureDirExists(filepath.Dir(filePath))
	return ioutil.WriteFile(filePath, contents, os.ModePerm)
}

//	A short-hand for `ioutil.WriteFile`, using `ModePerm`.
//	Also ensures the target file's directory exists.
func WriteTextFile(filePath, contents string) error {
	return WriteBinaryFile(filePath, []byte(contents))
}
