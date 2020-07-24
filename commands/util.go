package commands

import (
	"github.com/Plloi/pdb-cmdr/pkg/router"
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

func notImplemented(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "Command reserved, but not implemented")
}

// Setup Registers package commands to a command router
func Setup(r *router.CommandRouter) {
	log.Info("Setting up 'echo' command")
	r.RegisterCommand("echo", "Make the bot say something (Server Admin only)", echo)
	log.Info("Setting up '8ball' command")
	r.RegisterCommand("8ball", "Get A yes/answer from the magic B-ball.", Roll8Ball)
	// SAL := NewSAL()
	// r.RegisterCommand("sal", "Smash Amiibo League parent command", SAL.SAL)
}
