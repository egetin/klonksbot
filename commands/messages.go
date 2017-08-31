package commands

import (
	"fmt"
	"regexp"
	"strings"

	irc "github.com/fluffle/goirc/client"
)

func Parse(conn *irc.Conn, line *irc.Line) {
	// Parse incoming messages
	// Commands start with a dot (.)
	// Events (like tsips) are wrapped between * (*tsips*)

	fmt.Printf("%s | %s\n", line.Raw, line.Text())

	// Regex for matching command message
	command_re, err := regexp.Compile(`^(\.\w+).*$`)
	if err != nil {
		fmt.Println("Error while compiling command regexp: ", err.Error())
	}

	// Regex for catching events in message
	event_re, err := regexp.Compile(`\*\w+\*`)
	if err != nil {
		fmt.Println("Error while compiling event regexp: ", err.Error())
	}

	// Check for commands first, since an event won't exist in the same line.
	// ...well, it could, but we don't care :)
	// Check using regular expressions compiled above.
	if command_re.MatchString(line.Text()) {
		// Matched, split string into an array and check if command is valid
		fmt.Println("Someone entered a command")

		var cmd_arr []string = strings.Fields(line.Text())
		handleCommand(cmd_arr, conn, line)

	} else if event_re.MatchString(line.Text()) {
		// Matched, here we probably should do something with this
		fmt.Println("An event happened")
		events := event_re.FindAllString(line.Text(), -1)
		conn.Privmsg(line.Target(), fmt.Sprintf("You sent %d events at once!", len(events)))
	}
}

func handleCommand(cmd_arr []string, conn *irc.Conn, line *irc.Line) {
	// For testing purposes
	commands := []string{
		".op",
		".help",
		".tsips",
		".auth",
		".quit",
	}

	for _, command := range commands {
		if command == cmd_arr[0] {
			// Valid command, continue
			conn.Privmsg(line.Target(), fmt.Sprintf("Valid command!"))
			return
		}
	}

	conn.Privmsg(line.Target(), fmt.Sprintf("Invalid command!"))
}
