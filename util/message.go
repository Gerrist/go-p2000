package util

import (
	"strings"
	"time"
)

type Message = struct {
	received  time.Time
	text      string
	cap_codes []string
}

func MessageFromString(message string) Message {
	parts := strings.Split(message, "|")
	// 0 = FLEX
	// 1 = date (2023-08-28 20:43:07)
	// 2 = unknown
	// 3 = unknown
	// 4 = capcodes (000120901 000923993)
	// 5 = ALN
	// 6 = message

	// split parts[4] on spaces
	capCodes := strings.Split(parts[4], " ")
	// convert parts[1] to time.Time and put result in var called messageTime
	messageTime, _ := time.Parse("2006-01-02 15:04:05", parts[1])

	return Message{
		received:  messageTime,
		text:      parts[6],
		cap_codes: capCodes,
	}
}

func MessageDebugString(message Message) string {
	return "Message{received: " + message.received.String() + ", text: " + message.text + ", cap_codes: " + strings.Join(message.cap_codes, ",") + "}"
}
