package textcommand

import (
	"github.com/bwmarrin/discordgo"

	"github.com/sysnote8main/discordgo-bot-template/pkg/commands"
	"github.com/sysnote8main/discordgo-bot-template/pkg/util/timeutil"
)

func Register(cmdManager *commands.TextCommandManager) {
	cmdManager.AddCommand(commands.TextCommand{
		CommandName: "ping",
		Aliases:     []string{},
		Description: "respond pong!",
		HandlerFunc: func(s *discordgo.Session, e *discordgo.MessageCreate, args []string) error {
			_, err := s.ChannelMessageSend(e.ChannelID, "pong!")
			return err
		},
	})

	cmdManager.AddCommand(commands.TextCommand{
		CommandName: "help",
		Aliases:     []string{},
		Description: "help command",
		HandlerFunc: func(s *discordgo.Session, e *discordgo.MessageCreate, args []string) error {
			g := &discordgo.MessageEmbed{
				Title:       "Bot Help",
				Description: "List of commands",
				Fields:      cmdManager.GetHelpCommandFields(),
				Timestamp:   timeutil.GetNowTimeStamp(),
			}
			_, err := s.ChannelMessageSendEmbed(e.ChannelID, g)
			if err != nil {
				return err
			}
			return nil
		},
	})
}
