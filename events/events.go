package events

import (
	"fmt"
	"log"

	irc "github.com/fluffle/goirc/client"
)

type IrcEvent func(*irc.Line) (string, error)

const (
	EVENT_DECORATOR = "*"
)

var (
	events = map[string]IrcEvent{}
)

func registerEvent(eventName string, callback IrcEvent) {
	var parsedEventName string

	parsedEventName = parseEventName(eventName)

	if events[parsedEventName] != nil {
		// Event already exists, must not continue
		panic(fmt.Sprintf("Event %s already exists!\n", eventName))

	}

	events[parsedEventName] = callback
}

func parseEventName(eventName string) (parsedName string) {
	parsedName = EVENT_DECORATOR + eventName + EVENT_DECORATOR

	return
}

func HandleEvent(event string, conn *irc.Conn, line *irc.Line) {
	var err error
	var resp string

	eventFunc := events[event]
	if eventFunc == nil {
		// Invalid event
		conn.Privmsg(line.Target(), fmt.Sprintf("Invalid event!"))
		return
	}

	resp, err = eventFunc(line)
	if err != nil {
		log.Printf("User %s tried to do event %s but failed: %s\n", line.Nick, event, err)
		return
	}

	if resp != "" {
		conn.Privmsg(line.Target(), resp)
		return
	}

	return
}
