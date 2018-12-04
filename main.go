package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
)

func main() {
	// function which will calmly handle any runtime errors
	defer func() {
		err := recover()
		if err != nil {fmt.Println("Error:", err)}
	}()

	// first, we will get a token from token.txt
	authToken, err := ioutil.ReadFile("./token.txt")
	check(err)

	// attempt to instantiate our client using this token
	discord, err := discordgo.New("Bot" + string(authToken))
	check(err)

	// register our message handler
	discord.AddHandler(messageCreate)

	// open websocket connection
	err = discord.Open()
	check(err)
	defer discord.Close()

	// now our main thread simply waits for an interrupt
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Interrupt, os.Kill)
	<-sc
}

// messageCreate is our handler for messages
func messageCreate(sess *discordgo.Session, msg *discordgo.MessageCreate) {
	// ignore messages from self
	if msg.Author.ID == sess.State.User.ID {return}
}
