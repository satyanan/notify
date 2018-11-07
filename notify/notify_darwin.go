package notify

func notifyCmd(msg string) string {
	return "osascript -e 'tell Application \"Finder\" to display dialog \"" + msg + "\""
}
