package cmd

import (
	"path/filepath"
	"reflect"
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
		if testcase.srcIsExists == false {
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

var getToProjectsTests = []struct {
	name   string
	src    string
	expect []string
}{
	{
		name:   "should get git repositories for each user",
		src:    "test/.gitbucket",
		expect: []string{"backet.git", "backet.wiki.git"},
	},
}

func TestGetToProject(t *testing.T) {
	for _, testcase := range getToProjectsTests {
		t.Log(testcase.name)
		toProjects, err := getToProjects(testcase.src)
		for i, project := range toProjects {
			toProjects[i] = filepath.Base(project)
		}
		if err != nil {
			t.Error(err)
		}
		if !reflect.DeepEqual(toProjects, testcase.expect) {
			t.Errorf("=> Got %#v,\n => expected %#v", toProjects, testcase.expect)
		}
	}
}
