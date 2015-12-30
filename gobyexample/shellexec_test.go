package gobyexample

import (
	"testing"
)

func TestExecShellCmd(t *testing.T) {
	should_success := [][]string{
		[]string{"/usr/bin/docker", "images"},
		[]string{"/usr/bin/docker", "ps"},
		[]string{"/usr/bin/docker", "info"},
	}
	for _, command := range should_success {
		if output, err := ExecShellCmdNoPipe(command...); err != nil {
			t.Errorf("should success run \"%s\", but result in error: %s\n", command, err)
		} else {
			t.Logf("\"%s\": %s\n", command, string(output))
		}
	}
}
