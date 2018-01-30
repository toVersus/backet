package util

import (
	"path/filepath"
	"reflect"
	"testing"
)

var cmdStringTests = []struct {
	name    string
	command Command
	expect  string
}{
	{
		name:    "should return command without any arguments",
		command: Command{name: "cp"},
		expect:  "cp",
	},
	{
		name:    "should return command with any arguments",
		command: Command{name: "cp", args: []string{"-r", "-p"}},
		expect:  "cp -r -p",
	},
}

func TestString(t *testing.T) {
	t.Log("converting to string...")

	for _, testcase := range cmdStringTests {
		t.Log(testcase.name)
		if !reflect.DeepEqual(testcase.command.String(), testcase.expect) {
			t.Errorf("result => %#v\n  expected => %#v\n", testcase.command.String(), testcase.expect)
		}
	}
}

var cmdTests = []struct {
	name    string
	path    string
	command Command
}{
	{
		name:    "should not return any errors while executing command",
		command: Command{name: "dir"},
	},
}

func TestRunInDir(t *testing.T) {
	for _, testcase := range cmdTests {
		t.Log(testcase.name)
		pwd, _ := filepath.Abs(".")
		testcase.path = filepath.Join(filepath.Dir(pwd), "cmd")
		if _, err := testcase.command.RunInDir(testcase.path); err != nil {
			t.Error(err)
		}
	}
}
