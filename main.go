package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	irc "github.com/fluffle/goirc/client"

	"github.com/egetin/klonksbot/commands"
	"github.com/egetin/klonksbot/events"
)

// Configuration
// TODO: Parse configuration from a separate file
const (
	NICK     = "Klonkswagen"
	IDENT    = "klonkswagen"
	REALNAME = "Klonkswagen The Bot"
	SERVER   = "open.ircnet.net"
	PORT     = "6667"
	QUITMSG  = "Goodbye cruel world ;_;"
	CHANNEL  = "#egetestaa"
)

func main() {
	fmt.Println("Starting Klonkswagen")

	// Initialize configuration
	cfg := irc.NewConfig(NICK, IDENT, REALNAME)
	cfg.SSL = false
	cfg.Server = fmt.Sprintf("%s:%s", SERVER, PORT)
	cfg.NewNick = func(n string) string { return n + "_" }
	cfg.QuitMessage = QUITMSG

	// Connect to server
	c := irc.Client(cfg)
	quit := make(chan bool)
	addHandlers(c, quit)

	fmt.Printf("Connecting to %s:%s\n", SERVER, PORT)
	if err := c.Connect(); err != nil {
		fmt.Printf("Connection error: %s\n", err.Error())
	}

	<-quit
	fmt.Println("Quitting.")
	c.Close()
}

func addHandlers(c *irc.Conn, quit chan bool) {
	c.HandleFunc(irc.CONNECTED,
		func(conn *irc.Conn, line *irc.Line) {
			fmt.Println("Connected.")
			conn.Join(CHANNEL)
		})

	c.HandleFunc(irc.DISCONNECTED,
		func(conn *irc.Conn, line *irc.Line) {
			fmt.Println(line.Raw)
			quit <- true
		})

	c.HandleFunc(irc.PRIVMSG, parseMsg)
}

func parseMsg(conn *irc.Conn, line *irc.Line) {
	// Parse incoming messages

	// Regex for matching command message
	command_re, err := regexp.Compile(`^(\.\w+).*$`)
	if err != nil {
		log.Println("Error while compiling command regexp: ", err.Error())
	}

	// Regex for catching events in message
	event_re, err := regexp.Compile(`\*\w+\*`)
	if err != nil {
		log.Println("Error while compiling event regexp: ", err.Error())
	}

	// Check for commands and events using the regexes compiled above
	if command_re.MatchString(line.Text()) {
		// Matched, split string into an array and check if command is valid
		var cmd_arr []string = strings.Fields(line.Text())

		commands.HandleCommand(cmd_arr, conn, line)

	} else if event_re.MatchString(line.Text()) {
		var msgEvents []string = event_re.FindAllString(line.Text(), -1)

		for _, event := range msgEvents {
			events.HandleEvent(event, conn, line)
		}
	}
}
