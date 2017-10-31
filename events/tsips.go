package events

import (
	"fmt"

	irc "github.com/fluffle/goirc/client"
)

func init() {
	registerEvent("tsips", eventTsips)
}

func eventTsips(line *irc.Line) (resp string, err error) {
	resp = fmt.Sprintf("%s is drunk!!!!!\n", line.Nick)

	return
}
