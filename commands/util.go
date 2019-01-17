package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/plloi/Junior/router"
)

func notImplemented(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "Command reserved, but not implemented")
}

func Setup(r *router.CommandRouter) {
	r.RegisterCommand("echo", "Make the bot say something (Server Admin only)", echo)
	r.RegisterCommand("8ball", "Get A yes/answer from the magic B-ball.", Roll8Ball)
	r.RegisterCommand("sal", "Smash Amiibo League parent command", NewSAL().SAL)
	r.RegisterCommand("fortune", "Hopefully something whitty!", fortune)
	r.RegisterCommand("tf", "Hopefully something whitty!", timedFortune)
}
