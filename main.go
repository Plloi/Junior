package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Plloi/Junior/commands"
	"github.com/Plloi/pdb-cmdr/pkg/router"
	"github.com/Plloi/pdb-pokemon/pkg/pokemon"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

// Variables used for command line parameters
var (
	Token  string
	Router *router.CommandRouter
	SAL    *commands.SAL
)

func init() {

	flag.StringVar(&Token, "t", "", "Bot Discord Token")
	flag.Parse()
	log.SetLevel(log.DebugLevel)
	godotenv.Load()

	if Token == "" {
		Token = os.Getenv("DISCORD_TOKEN")
	}

}

func main() {
	log.Info("Starting up")
	// Create a new Discord session using the provided bot token.
	if Token == "" {
		log.Error("Token Required for bot usage")
		os.Exit(1)
	}
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	Router = router.NewCommandRouter()
	Router.DefaultPrefix = "pj!"
	Router.RegisterCommand("prefix", "Sets the bot command prefix (Admin Locked)", Router.SetPrefix)

	// Register router's command handler for message events.
	dg.AddHandler(Router.HandleCommand)

	log.Info("Importing commands module")
	commands.Setup(Router)

	log.Info("Loading Pokemon Module")
	pokemon.Setup(Router)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	log.Info("Bot is now running")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}
