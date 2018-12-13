package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/plloi/Junior/router"
)

func echo(s *discordgo.Session, m *discordgo.MessageCreate) {
	if ok, _ := router.MemberHasPermission(s, m.GuildID, m.Author.ID, discordgo.PermissionAdministrator); ok {
		s.ChannelMessageSend(m.ChannelID, m.Content)
	}
}
