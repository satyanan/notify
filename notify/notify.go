package notify

import (
	"fmt"
	"os/exec"
	"strings"
)

// Notify sends a notification to system
func Notify(cmd string) {
	execCmd(notifyCmd)
}

func execCmd(cmd string) {
	fmt.Println("command is ", cmd)
	parts := strings.Fields(cmd)
	command := parts[0]
	args := parts[1:len(parts)]

	out, err := exec.Command(command, args...).Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Printf("Out %s", out)
}
