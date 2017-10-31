package commands

import irc "github.com/fluffle/goirc/client"

func init() {
	registerCommand("op", commandOp)
}

func commandOp(args []string, line *irc.Line) (resp string, err error) {
	return
}
