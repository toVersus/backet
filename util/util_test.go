package util

import (
	"os"
	"path/filepath"
	"testing"
)

var dirTests = []struct {
	name     string
	path     string
	filename string
}{
	{
		name: "create and delete single directory",
		path: "foo",
	},
	{
		name: "create and delete directory containing child directory",
		path: "foo/bar/baz",
	},
}

func TestCreateDeleteDir(t *testing.T) {
	for _, testcase := range dirTests {
		t.Log(testcase.name)
		if err := CreateDir(testcase.path); err != nil {
			t.Error(err)
		}
		if testcase.filename != "" {
			path := filepath.Join(testcase.path, testcase.filename)
			if _, err := os.Create(path); err != nil {
				t.Error(err)
			}
			if err := os.Chmod(path, 0444); err != nil {
				t.Error(err)
			}
		}
		if err := DeleteDir(GetParentDirName(testcase.path)); err != nil {
			t.Error(err)
		}
	}

}
