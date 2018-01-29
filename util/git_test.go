package util

import (
	"reflect"
	"testing"
)

var commandTest = []struct {
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

	for _, testcase := range commandTest {
		t.Log(testcase.name)
		if !reflect.DeepEqual(testcase.command.String(), testcase.expect) {
			t.Errorf("result => %#v\n  expected => %#v\n", testcase.command.String(), testcase.expect)
		}
	}
}
