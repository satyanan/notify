package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/satyanan/notify/notify"
)

const usageMessage = "" +
	`Usage of 'notify':
Given a long running command for ex.
	sleep(2)
Generate a notification when the command finishes
	notify sleep(2)
`

func usage() {
	fmt.Fprintln(os.Stderr, usageMessage)
	os.Exit(2)
}

var (
	message = flag.String("m", "", "Custom message for notification")
)

func main() {
	flag.Usage = usage
	flag.Parse()

	// Usage info when no flag or arg passed
	if flag.NFlag() == 0 && flag.NArg() == 0 {
		flag.Usage()
	}

	if *message == "" {
		fmt.Println(flag.Args())
		var cmd string
		for _, str := range flag.Args() {
			cmd = cmd + str
		}
		notify.Notify(cmd)
		// message = flag.Args()
	}
}
