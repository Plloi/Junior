package commands

import (
	"github.com/Plloi/pdb-cmdr/pkg/router"
	"github.com/bwmarrin/discordgo"
)

func echo(s *discordgo.Session, m *discordgo.MessageCreate) {
	if ok, _ := router.MemberHasPermission(s, m.GuildID, m.Author.ID, discordgo.PermissionAdministrator); ok {
		s.ChannelMessageSend(m.ChannelID, m.Content)
	}
}
