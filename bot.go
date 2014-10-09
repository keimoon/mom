package main

import (
	"github.com/daneharrigan/hipchat"
	"github.com/keimoon/cerebro"
	"log"
	"strings"
)

type Bot struct {
	user    string
	name    string
	mention string
	pass    string
	client  *hipchat.Client
}

func NewBot(user, name, mention, pass string) (*Bot, error) {
	b := &Bot{user: user, name: name, mention: mention, pass: pass}
	client, err := hipchat.NewClient(user, pass, "bot")
	if err != nil {
		return nil, err
	}
	b.client = client
	client.Status("chat")
	return b, nil
}

func (b *Bot) Join(room string) {
	roomJid := "" + room + "@conf.hipchat.com"
	b.client.Join(roomJid, b.name)
}

func (b *Bot) Say(room string, msg string) {
	roomJid := "" + room + "@conf.hipchat.com"
	b.client.Say(roomJid, b.name, msg)
}

func (b *Bot) Reply(msg string) string {
        msg = msg[len("@"+b.mention)+1:]
	answer, _ := cerebro.Cerebro.Ask(msg)	
        return answer
}

func (b *Bot) Listen(room string) {
	go b.client.KeepAlive()

	for message := range b.client.Messages() {
		log.Printf("%s: %s\n", message.From, message.Body)
		input := message.Body
		if strings.HasPrefix(input, "@"+b.mention) {
			output := b.Reply(input)
			b.Say(room, output)
		}
	}
}
