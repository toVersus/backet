package cmd

import (
	"testing"

	"github.com/toversus/backet/util"
)

var gitCloneTests = []struct {
	name        string
	src         string
	srcIsExists bool
	dest        string
}{
	{
		name:        "should skip git cloning due to existing dest directory",
		src:         "test/.gitbucket/repositories/toversus/backet.git",
		srcIsExists: true,
		dest:        "test/backup",
	},
}

func TestCreateMirrorRepository(t *testing.T) {
	for _, testcase := range gitCloneTests {
		t.Log(testcase.name)
		if testcase.srcIsExists == true {
			_ = util.CreateDir(testcase.src)
		}
		if err := createMirrorRepository(testcase.src, testcase.dest); err != nil {
			t.Error(err)
		}
		if err := util.DeleteDir(testcase.dest); err != nil {
			t.Error(err)
		}
	}
}
