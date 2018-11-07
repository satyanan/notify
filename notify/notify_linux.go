package notify

func notifyCmd(msg string) string {
	return "notify-send " + msg
}
