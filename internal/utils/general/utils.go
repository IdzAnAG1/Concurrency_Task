package general

import (
	"crypto/md5"
	"encoding/hex"
	"os"
)

func ConvertToHash(content string) string {
	hasher := md5.New()
	hasher.Write([]byte(content))
	return hex.EncodeToString(hasher.Sum(nil))
}

func GetFileContents(pathToDir string, fileName string) (string, error) {
	file, err := os.OpenFile(pathToDir+"/"+fileName, os.O_RDONLY, os.ModePerm)
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
