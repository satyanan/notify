package notify

func notifyCmd(msg string) string {
	return "toast -t \"notify\" -m \"" + msg + "\""
}
