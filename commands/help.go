package commands

import (
	"fmt"

	irc "github.com/fluffle/goirc/client"
)

func init() {
	registerCommand("help", commandHelp)
}

func commandHelp(args []string, line *irc.Line) (resp string, err error) {
	// TODO: After moving commands to a clever place, use it here too!
	resp = fmt.Sprintf("Available commands: .op .help .tsips\n")

	return
}
