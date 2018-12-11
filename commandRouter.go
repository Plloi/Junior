package main

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

//CommandRouter Handles routing of chat commands to handler functions
type CommandRouter struct {
	commands      map[string]func(*discordgo.Session, *discordgo.MessageCreate)
	helpText      map[string]string
	commandPrefix string
}

// NewCommandRouter Read the name
func NewCommandRouter() *CommandRouter {
	// cMap := make(map[string]func(*discordgo.Session, *discordgo.MessageCreate))
	router := &CommandRouter{
		commands: make(map[string]func(*discordgo.Session, *discordgo.MessageCreate)),
		helpText: make(map[string]string),
	}
	router.commandPrefix = "sal!"

	router.setup()

	return router
}

func (c *CommandRouter) setup() {
	c.commands["help"] = c.help
	c.helpText["help"] = "This help text"
}

func (c *CommandRouter) help(s *discordgo.Session, m *discordgo.MessageCreate) {
	var helpMessage = "Here's what I can do!\n"
	for key, value := range c.helpText {
		helpMessage += fmt.Sprintf("* %s: %s\n", key, value)
	}
	s.ChannelMessageSend(m.ChannelID, helpMessage)
}

// RegisterCommand Adds a command, it's help text, and function to the router
func (c *CommandRouter) RegisterCommand(command string, help string, f func(*discordgo.Session, *discordgo.MessageCreate)) error {
	if _, ok := c.commands[command]; ok {
		return fmt.Errorf("Command %s is already registered", command)
	}
	if _, ok := c.helpText[command]; ok {
		return fmt.Errorf("Help for command %s is already registered", command)
	}

	c.commands[command] = f
	c.helpText[command] = help
	return nil
}

// HandleCommand Takes Discord input and tries to find a relevant command, can be passed to discord-go's AddHandler
func (c *CommandRouter) HandleCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Check for bot trigger
	if len(m.Content) > len(c.commandPrefix) && m.Content[:len(c.commandPrefix)] == c.commandPrefix {
		// trim prefix
		m.Content = m.Content[len(c.commandPrefix):]
		args := strings.Split(m.Content, " ")

		// Check if command is registered
		if f, ok := c.commands[args[0]]; ok {
			// Remove command
			m.Content = m.Content[len(args[0]):]
			// Call function
			f(s, m)
		} else {
			s.ChannelMessageSend(m.ChannelID, "Command not recognized")
		}
	}
}
