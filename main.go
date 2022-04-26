package main

import (
	"flag" // This package implements commandline flag parsing (the '-option_name' options for a CLI program).
	"fmt"
	"github.com/bwmarrin/discordgo"
)

// "var" keyword is used to create a variable of a particular type having a proper name and proper value.
var (
	Token string
)

// The "init()" function is used to set up some form of state on the initial startup of the program.
// In this case, it gets and stores the Discord bot token.
func init() {
	flag.StringVar(&Token, "token", "", "Bot private token")
	flag.Parse()
}

func main() {
	discordgo, err := discordgo.New("Bot " + Token) // Create a new Discord session using the provided bot token.

	// If there is an error, Print a message with the error and stop the program,
	if err != nil {
		fmt.Println("There was an error while trying to connect to Discord Bot,", err)
		return
	}
	discordgo.AddHandler(messageCreate) // Creates an event handler, receiving the messageCreate function.
}

// messageCreat function: receives the discordgo.Session (the bot) and the MessageCreate method
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content != "ping" {
		return
	}
}
