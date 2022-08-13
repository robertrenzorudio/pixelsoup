package runcommand

import (
	"bytes"
	"errors"
	"os/exec"

	"github.com/google/shlex"
)

func Run(command string) (stdout *bytes.Buffer, stderr *bytes.Buffer, err error) {
	args, err := shlex.Split(command)

	if err != nil {
		return nil, nil, err
	}

	cmd := exec.Command(args[0], args[1:]...)
	stdout, stderr = &bytes.Buffer{}, &bytes.Buffer{}

	cmd.Stdout = stdout
	cmd.Stderr = stderr

	err = cmd.Run()

	if err != nil {
		return nil,
			stderr,
			err
	}

	if stderr.Len() != 0 {
		return nil,
			stderr,
			errors.New("an error occurred")
	}

	return stdout, nil, nil
}
