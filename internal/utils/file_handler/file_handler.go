package file_handler

import (
	"log"
	"os"
	"path"
)

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
