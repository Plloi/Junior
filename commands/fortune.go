package commands

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/plloi/Junior/router"
)

func fortune(s *discordgo.Session, m *discordgo.MessageCreate) {
	out, err := exec.Command("fortune", "-s").Output()
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "A greivous error hass occured")
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Error: %v", err))
		return
	}
	s.ChannelMessageSend(m.ChannelID, string(out))

}

func timedFortune(s *discordgo.Session, m *discordgo.MessageCreate) {
	if ok, _ := router.MemberHasPermission(s, m.GuildID, m.Author.ID, discordgo.PermissionAdministrator); !ok {
		return
	}

	ticker := time.NewTicker(30 * time.Second)
	go func(s *discordgo.Session, m *discordgo.MessageCreate) {
		for {
			select {
			case <-ticker.C:
				fortune(s, m)
			default:
				time.Sleep(3 * time.Second)
			}
		}
	}(s, m)

}
