package bot

import (
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"

	"github.com/sysnote8main/discordgo-bot-template/pkg/commands"
	"github.com/sysnote8main/discordgo-bot-template/pkg/commands/textcommand"
	"github.com/sysnote8main/discordgo-bot-template/pkg/config"
)

var (
	textCmdManager *commands.TextCommandManager
)

func init() {
	config.ReadConfig()

	textCmdManager = commands.NewTextCommandManager(config.GetConfig().Prefix)
	textcommand.Register(textCmdManager)
}

func Run() {
	fmt.Println("Registered command keys:", textCmdManager.GetAllCommandKeys())
	dg, err := discordgo.New("Bot " + config.GetConfig().Token)
	if err != nil {
		slog.Error("Failed to create discord session", slog.Any("error", err))
		os.Exit(1)
	}
	dg.Identify.Intents = discordgo.IntentsGuilds + discordgo.IntentsMessageContent + discordgo.IntentsGuildMessages

	dg.AddHandler(onReady)
	dg.AddHandler(textCmdManager.OnEvent)

	if err = dg.Open(); err != nil {
		slog.Error("Failed to open connection", slog.Any("error", err))
		os.Exit(1)
	}

	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-s

	slog.Info("Shutting down...")
	err = dg.Close()
	if err != nil {
		slog.Error("Failed to close connection", slog.Any("error", err))
	}
	slog.Info("Shutdown completed.")
	// TODO save config which edit from discord command
}

func onReady(s *discordgo.Session, _ *discordgo.Ready) {
	slog.Info(fmt.Sprintf("Bot on ready! (name:%s)", s.State.User.Username))
}
