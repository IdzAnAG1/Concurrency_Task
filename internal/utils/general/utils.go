package general

import (
	"crypto/md5"
	"encoding/hex"
	"log"
	"os"
	"path"
	"path/filepath"
)

func ConvertToHash(content string) string {
	hasher := md5.New()
	hasher.Write([]byte(content))
	return hex.EncodeToString(hasher.Sum(nil))
}

func GetFileContents(pathToDir string, fileName string) (string, error) {
	file, err := os.OpenFile(filepath.Join(pathToDir, fileName), os.O_RDONLY, os.ModePerm)
	if err != nil {
		return "", err
	}
	defer file.Close()

	buff := make([]byte, 1024)
	_, err = file.Read(buff)

	if err != nil {
		return "", err
	}
	return string(buff), nil
}

func ReadFromFile(dir string, file os.DirEntry) string {
	data, err := os.ReadFile(path.Join(dir, file.Name()))
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

func GetFilesInDirectory(path string) []os.DirEntry {
	FilesIntoDirectory, err := os.ReadDir(path + "/")
	if err != nil {
		log.Fatal(err)
	}
	return FilesIntoDirectory
}
