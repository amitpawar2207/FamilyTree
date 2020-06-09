package fileoperations

import (
	"fmt"
	"reflect"
	"testing"
)

type result struct {
	inputs []string
	err    error
}

func TestReadInputFile(t *testing.T) {
	testData := []struct {
		testName   string
		filePath   string
		resultData result
	}{
		{
			"File doesn't exists",
			"./resources/test_No_File.txt",
			result{
				inputs: make([]string, 0),
				err:    fmt.Errorf("no such file or directory exists"),
			},
		},
		{
			"Empty File",
			"./resources/test_family_members_empty.txt",
			result{
				inputs: make([]string, 0),
				err:    fmt.Errorf("file is empty"),
			},
		},
		{
			"valid File",
			"./resources/test_family_members_valid.txt",
			result{
				inputs: []string{
					"ADD_PARTNER Vich Lika Female",
					"ADD_CHILD_FOR_PREDEFINED_TREE Anga Aras Male",
					"ADD_PARTNER Aras Chitra Female",
					"ADD_CHILD_FOR_PREDEFINED_TREE Anga Satya Female",
					"ADD_PARTNER Satya Vyan Male"},
				err: nil,
			},
		},
	}

	for _, td := range testData {
		t.Run(td.testName, func(t *testing.T) {
			inputs, err := ReadInputFile(td.filePath)
			var r result
			r.err = err
			r.inputs = inputs
			if reflect.DeepEqual(r, td.resultData) {
				t.Errorf("Expected output is %v for test case %v and got %v", td.resultData.err, td.testName, err)
			}

		})
	}

}
