package commands

import "github.com/bwmarrin/discordgo"

func notImplemented(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "Command reserved, but not implemented")
}
