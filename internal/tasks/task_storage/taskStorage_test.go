package task_storage

import (
	"fmt"
	"testing"
)

type test struct {
	Name    string
	wantErr bool
}

func (te *test) Launch() {
	fmt.Printf("Test File | Test \"ConcurrencyTask\" implementation | Structure name : %s\n", te.Name)
}
func TestTaskStorage_AddInStorage_NewTaskStorage(t *testing.T) {
	Store := NewTaskStorage()
	if Store == nil {
		t.Errorf("Storage Object is not created ")
	}

	tests := []test{
		{Name: "wer", wantErr: false},
		{Name: "te", wantErr: false},
		{Name: "test_13", wantErr: false},
		{Name: "test_123", wantErr: false},
		{Name: "", wantErr: true},
	}
	for i := 0; i < len(tests); i++ {
		Store.AddInStorage(tests[i].Name, &test{Name: tests[i].Name})
	}

	if len(Store.GetKeys()) != len(tests) {
		t.Errorf("Something was wrong keys : %v and exp : %v", len(Store.GetKeys()), len(tests))
	}
	for _, el := range tests {
		if Store.taskIsLocatedInTheRepository(el.Name) == el.wantErr {
			t.Errorf("The task was not added to the storage : %s", el.Name)
		}
	}
}

func TestGetStorageInstance(t *testing.T) {
	type ttest struct {
		store   *TaskStorage
		wantErr bool
	}
	tests := []ttest{
		{store: NewTaskStorage(), wantErr: false},
		{store: nil, wantErr: true},
	}
	for _, el := range tests {
		if el.store != GetStorageInstance() && el.wantErr != false {
			t.Errorf("Error when getting a link to the repository")
		}
	}

}
