package commands

import (
	"fmt"
	"log"

	irc "github.com/fluffle/goirc/client"
)

type IrcCommand func([]string, *irc.Line) (string, error)

const (
	COMMAND_PREFIX = "."
)

var (
	commands = map[string]IrcCommand{}
)

func registerCommand(commandName string, callback IrcCommand) {
	var parsedCommandName string

	parsedCommandName = parseCommandName(commandName)

	if commands[parsedCommandName] != nil {
		// Command exists, must not continue
		panic(fmt.Sprintf("Command %s already exists in command pool!\n", commandName))
	}

	commands[parsedCommandName] = callback
}

func parseCommandName(commandName string) (parsedCommandName string) {
	parsedCommandName = COMMAND_PREFIX + commandName

	return
}

func HandleCommand(cmd_arr []string, conn *irc.Conn, line *irc.Line) {
	var err error
	var resp string

	cmd := commands[cmd_arr[0]]
	if cmd == nil {
		// Invalid command
		conn.Privmsg(line.Target(), fmt.Sprintf("Invalid command!"))
		return
	}

	args := cmd_arr[1:]
	resp, err = cmd(args, line)
	if err != nil {
		log.Printf("Command %s for user %s returned error: %s\n", cmd_arr[0], line.Nick, err)
		return
	}

	if resp != "" {
		conn.Privmsg(line.Target(), resp)
		return
	}
}
