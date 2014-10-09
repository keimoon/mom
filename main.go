package main

import (
	"log"
)

func main() {

	user := ""
	pass := ""
	name := ""
	mention := ""
	room := ""
	bot, err := NewBot(user, name, mention, pass)
	if err != nil {
		log.Fatal(err)
	}
	bot.Join(room)
	bot.Say(room, "Im home, sons")
	bot.Listen(room)

}
