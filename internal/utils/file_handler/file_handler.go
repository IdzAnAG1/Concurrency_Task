package file_handler

import (
	"os"
	"path"
)

func ReadFromFile(dir string, file os.DirEntry) (string, error) {
	data, err := os.ReadFile(path.Join(dir, file.Name()))
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func GetFilesInDirectory(path string) ([]os.DirEntry, error) {
	FilesIntoDirectory, err := os.ReadDir(path)
	if err != nil {
		return []os.DirEntry{}, err
	}
	return FilesIntoDirectory, nil
}
