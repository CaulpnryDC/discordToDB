package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/andersfylling/disgord"
	"github.com/hashicorp/consul/api"
)

func printMessage(session disgord.Session, evt *disgord.MessageCreate) {
	msg := evt.Message
	fmt.Println(msg.Author.String() + ": " + msg.Content)
	log.Println(msg.Author.String() + ": " + msg.Content)
}

func main() {
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}

	kv := client.KV()

	pair, _, err := kv.Get("discord/starfinder_looter/bot_token", nil)
	if err != nil {
		panic(err)
	}

	//token := string(pair.Value)

	dClient := disgord.New(disgord.Config{
		BotToken: "token",
	})

	defer dClient.StayConnectedUntilInterrupted(context.Background())

	t := time.Now()
	ts := t.String()
	update := WriteToFile("/app/files/loot.txt", ts)

	dClient.On(disgord.EvtMessageCreate, printMessage)
	log.Println(update, printMessage)
}
