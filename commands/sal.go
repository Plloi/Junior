package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/plloi/Junior/database"
	"github.com/plloi/Junior/models"
	"github.com/plloi/Junior/router"
)

// SAL Struct for the Smash Amiibo League Bot Module
type SAL struct {
	// Amiibos Known Amiibo figures
	Amiibos []models.Amiibo
	// Amiibo Figures configured as figters in the league
	Fighters []models.Fighter
	// Upcoming matches
	Matches []models.Match

	// router internal Command router for sub commmands
	router *router.CommandRouter
}

// NewSAL Return a new SAL
func NewSAL() *SAL {
	sal := &SAL{}
	sal.getAmiibo()
	sal.getFighters()
	sal.getMatches()
	sal.router = router.NewCommandRouter()
	sal.router.CommandPrefix = ""
	sal.router.RegisterCommand("add", "Adds a new Amiibo or figther (Not implemented)", sal.add)
	return sal
}

func (sal *SAL) getAmiibo() {
	//TODO: fetch from datastore
	sal.Amiibos = make([]models.Amiibo, 0)
	return
}
func (sal *SAL) getFighters() {
	//TODO: fetch from datastore
	sal.Fighters = make([]models.Fighter, 0)
	return
}
func (sal *SAL) getMatches() {
	//TODO: fetch from datastore
	sal.Matches = make([]models.Match, 0)
	return
}

//SAL Base command for the sal module
func (sal *SAL) SAL(s *discordgo.Session, m *discordgo.MessageCreate) {
	sal.router.HandleCommand(s, m)
}
func (sal *SAL) add(s *discordgo.Session, m *discordgo.MessageCreate) {
	DB := database.NewDatastore("Meep")
	Test := &models.Amiibo{
		Serial:      "10000",
		Description: "Test Amiibo",
	}
	id, err := DB.Add(Test)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Error: %v", err))
	} else {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("ID: %v", id))
	}
	// notImplemented(s, m)
	return
}
