package handler_tcp_transport

import "fmt"

func help() string {
	return fmt.Sprint(
		"Commands:\n" +
			"  /register <username> <password> — create account\n" +
			"  /login <username> <password>    — authenticate\n" +
			"  /rooms                          — list available rooms\n" +
			"  /join <room>                    — join or create a room\n" +
			"  /msg <text>                     — send message to current room\n" +
			"  /history                        — last 50 messages in current room\n",
	)
}
