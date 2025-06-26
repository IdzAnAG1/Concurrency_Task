package general

import (
	"crypto/md5"
	"encoding/hex"
	"log"
	"os"
	"path"
)

func ConvertToHash(content string) string {
	hasher := md5.New()
	hasher.Write([]byte(content))
	return hex.EncodeToString(hasher.Sum(nil))
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
