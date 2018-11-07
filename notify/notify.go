package notify

import (
	"os/exec"
	"strings"
)

// Notify sends a notification to system
func Notify(cmd string, msg string) error {
	err := execCmd(cmd)
	execCmd(notifyCmd(msg))
	if err != nil {
		return err
	}
	return nil

}

func execCmd(cmd string) error {
	parts := strings.Fields(cmd)
	command := parts[0]
	args := parts[1:len(parts)]

	_, err := exec.Command(command, args...).Output()
	if err != nil {
		return err
	}

	return nil
}
