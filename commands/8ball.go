package commands

import (
	"math/rand"
	"time"

	"github.com/bwmarrin/discordgo"
)

var standardAnswers = [20]string{
	"It is certain.",
	"It is decidedly so.",
	"Without a doubt.",
	"Yes - definitely.",
	"You may rely on it.",
	"As I see it, yes.",
	"Most likely.",
	"Outlook good.",
	"Yes.",
	"Signs point to yes.",
	"Reply hazy, try again.",
	"Ask again later.",
	"Better not tell you now.",
	"Cannot predict now.",
	"Concentrate and ask again.",
	"Don't count on it.",
	"My reply is no.",
	"My sources say no.",
	"Outlook not so good.",
	"Very doubtful.",
}

// Roll8Ball Rolls a standard Magic 8-ball
func Roll8Ball(s *discordgo.Session, m *discordgo.MessageCreate) {
	rand.Seed(time.Now().UnixNano())
	s.ChannelMessageSend(m.ChannelID, standardAnswers[rand.Intn(len(standardAnswers))])
}
