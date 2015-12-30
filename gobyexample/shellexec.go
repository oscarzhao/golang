package gobyexample

import (
	"errors"
	"io/ioutil"
	"os/exec"
)

// report all error information to error
func ExecShellCmdNoPipe(commands ...string) ([]byte, error) {
	cmd := exec.Command(commands[0], commands[1:]...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	//cmd error
	stderr, errPipe := cmd.StderrPipe()
	if errPipe != nil {
		return nil, errPipe
	}

	if err := cmd.Start(); err != nil {
		return nil, err
	}

	bytesErr, err := ioutil.ReadAll(stderr)
	if err != nil {
		return nil, err
	}

	if len(bytesErr) != 0 {
		return nil, errors.New(string(bytesErr))
	}

	bytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		return nil, err
	}

	if err := cmd.Wait(); err != nil {
		return nil, err
	}
	if len(bytes) != 0 {
		return bytes, nil
	}
	return nil, errors.New("no output")
}
