package file_handler

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestGetFilesInDirectory(t *testing.T) {
	type test struct {
		countFiles int
	}
	tests := []test{
		{countFiles: 14},
		{countFiles: 10},
		{countFiles: 11},
		{countFiles: 12},
		{countFiles: 13},
		{countFiles: 0},
	}
	for _, el := range tests {
		dir := t.TempDir()
		for i := 0; i < el.countFiles; i++ {
			_, errCreationFile := os.CreateTemp(dir, fmt.Sprintf("test_file_temp_%d_*.go", i))
			if errCreationFile != nil {
				t.Errorf("Error at creation file :%+v", errCreationFile)
			}
		}
		files, err := GetFilesInDirectory(dir)
		if err != nil {
			t.Errorf("Error at getting file from directory :%v", err)
		}
		t.Logf("Len Files :%d", len(files))
		if len(files) != el.countFiles {
			t.Errorf("[UB] The number of files read does not match the specified ones")
		}
	}

}

func TestReadFromFile(t *testing.T) {
	type test struct {
		textForWritingInFile string
		wantErr              bool
	}
	tests := []test{
		{textForWritingInFile: "123412341234\n123412341234", wantErr: false},
		{textForWritingInFile: "ABCD", wantErr: false},
		{textForWritingInFile: "package main\nfunc main() {}", wantErr: false},
		{textForWritingInFile: "func main(){}\nfunc int() {}", wantErr: false},
		{textForWritingInFile: "Test\nTest\nTest\nTest", wantErr: false},
		{textForWritingInFile: "   ", wantErr: false},
		{textForWritingInFile: " ", wantErr: false},
	}
	dir := t.TempDir()
	filesMap := make(map[string]string)
	for _, el := range tests {
		if file, err := os.CreateTemp(dir, "temp_test_*.go"); err == nil {
			if _, errWrite := file.WriteString(el.textForWritingInFile); errWrite != nil {
				t.Logf("[ERROR] - %v", errWrite)
			}
			filesMap[filepath.Base(file.Name())] = el.textForWritingInFile
			file.Close()
		} else {
			t.Logf("[ERROR] - %v", err)
			file.Close()
		}
	}
	for key, value := range filesMap {
		t.Logf("!T!T! - Key : %s ; Val : %s", key, value)
	}
	files, errGettingFiles := GetFilesInDirectory(dir)
	if errGettingFiles != nil {
		t.Logf("[ERROR] - %v", errGettingFiles)
	}
	for _, el := range files {
		if content, err := ReadFromFile(dir, el); err == nil {
			t.Logf("FILENAME:%s - [content %s ; exp %s", el.Name(), content, filesMap[el.Name()])
			if content != filesMap[el.Name()] {
				t.Errorf("[UB] file content is unexpected")
			}
		} else {
			t.Logf("[ERROR] - %v", err)
		}
	}
}
