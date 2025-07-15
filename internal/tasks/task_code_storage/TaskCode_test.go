package task_code_storage

import (
	"fmt"
	"go.uber.org/zap"
	"sync"
	"testing"
)

func getTestCodeStorage() *TCStorage {
	return &TCStorage{
		logger:  *zap.NewNop(),
		mu:      sync.Mutex{},
		Storage: make(map[string]string),
	}
}

// [UB] Unexpected Behavior

func TestTCStorage_Initialize(t *testing.T) {
	store := getTestCodeStorage()
	type test struct {
		pathToDir string
		wantErr   bool
	}
	tests := []test{
		{pathToDir: "", wantErr: true},
		{pathToDir: "./filename.ext", wantErr: true},
		{pathToDir: "./EmptyFolder", wantErr: true},
		{pathToDir: "./Test_NonEmptyFolder", wantErr: false},
	}
	for i, el := range tests {
		if err := store.Initialize(el.pathToDir); (err != nil) != el.wantErr {
			t.Errorf("[UB] Unexpected error : '%v' test number : %d", err, i)
		}
	}
}

func TestTCStorage_PutGetDelete_inConcurrency(t *testing.T) {

	type test struct {
		filename    string
		code        string
		wantToPlace bool
		getPointer  bool
	}

	store := getTestCodeStorage()

	testsForPut := []test{
		{filename: "test_1.go", code: "package tests", wantToPlace: true, getPointer: true},
		{filename: ".go", code: "package tests", wantToPlace: true},
		{filename: "t.go", code: "", wantToPlace: true},
		{filename: "test_1.go", code: "package test\nfunc test() {}", wantToPlace: true},
		{filename: "", code: "package tests", wantToPlace: false},
	}
	twg := sync.WaitGroup{}

	for _, el := range testsForPut {
		twg.Add(1)
		go func(te test) {
			defer twg.Done()
			err := store.Put(te.filename, te.code)
			if err != nil {
				t.Logf("%v", err)
			}
		}(el)
	}
	twg.Wait()
	for _, el := range store.GetKeys() {
		fmt.Println(el)
	}
	twg2 := sync.WaitGroup{}
	for _, el := range testsForPut {
		twg2.Add(1)
		go func(te test) {
			defer twg2.Done()
			if te.getPointer {
				return
			}
			code, err := store.Get(te.filename)
			if err != nil {
				t.Logf("%v", err)
				return
			}
			if code != te.code {
				t.Logf("getted string : '%s'; expected string '%s'", code, te.code)
				t.Errorf("[UB] The content doesn't add up in %s", te.filename)
			}
		}(el)
	}

	twg2.Wait()

	for _, el := range testsForPut {
		store.Delete(el.filename)
		if store.Len() == 0 {
			break
		}
	}
}
