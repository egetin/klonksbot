package main

import (
	"fmt"

	irc "github.com/fluffle/goirc/client"

	"github.com/egetin/klonksbot/commands"
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
	AddHandlers(c, quit)

	fmt.Printf("Connecting to %s:%s\n", SERVER, PORT)
	if err := c.Connect(); err != nil {
		fmt.Printf("Connection error: %s\n", err.Error())
	}

	<-quit
	fmt.Println("Quitting.")
	c.Close()
}

func AddHandlers(c *irc.Conn, quit chan bool) {
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

	c.HandleFunc(irc.PRIVMSG, commands.Parse)
}
