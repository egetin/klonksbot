package commands

import (
	"fmt"

	irc "github.com/fluffle/goirc/client"
)

func init() {
	registerCommand("tsips", commandTsips)
}

func commandTsips(args []string, line *irc.Line) (resp string, err error) {
	var target string

	if len(args) == 0 {
		target = line.Nick
	} else {
		target = args[0]
	}

	// TODO: Fetch statistics and parse them to a simple line!

	resp = fmt.Sprintf("Target is: %s\n", target)
	return
}
