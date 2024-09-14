package bot

import (
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"

	"github.com/sysnote8main/discordgo-bot-template/pkg/commands"
	"github.com/sysnote8main/discordgo-bot-template/pkg/commands/slashcommand"
	"github.com/sysnote8main/discordgo-bot-template/pkg/commands/textcommand"
	"github.com/sysnote8main/discordgo-bot-template/pkg/config"
)

var (
	textCmdManager         *commands.TextCommandManager
	slashCmdManager        *commands.SlashCommandManager
	removeSlashCommands    = false
	removeAllSlashCommands = true
)

func init() {
	// read configurations
	config.ReadConfig()

	// init variables
	textCmdManager = commands.NewTextCommandManager(config.GetConfig().Prefix)
	slashCmdManager = commands.NewSlashCommandManager()

	// Register commands
	textcommand.Register(textCmdManager)
	slashcommand.Register(slashCmdManager)
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
	dg.AddHandler(slashCmdManager.OnEvent)

	if err = dg.Open(); err != nil {
		slog.Error("Failed to open connection", slog.Any("error", err))
		os.Exit(1)
	}

	// Register slash commands to discord
	slashCmdManager.Register(dg, config.GetConfig().GuildId)

	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-s

	slog.Info("Shutting down...")

	// remove slash commands
	if removeSlashCommands {
		if removeAllSlashCommands {
			slog.Info("Removing all slash commands...")
		}
	}

	// close discord connection
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
