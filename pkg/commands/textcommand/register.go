package textcommand

import (
	"fmt"

	"github.com/bwmarrin/discordgo"

	"github.com/sysnote8main/discordgo-bot-template/pkg/commands"
)

func Register(cmdManager *commands.TextCommandManager) {
	cmdManager.AddCommand(commands.TextCommand{
		CommandName: "ping",
		Aliases:     []string{},
		Description: "respond pong!",
		HandlerFunc: func(s *discordgo.Session, e *discordgo.MessageCreate, args []string) {
			s.ChannelMessageSend(e.ChannelID, "pong!")
		},
	})

	cmdManager.AddCommand(commands.TextCommand{
		CommandName: "help",
		Aliases:     []string{},
		Description: "help command",
		HandlerFunc: func(s *discordgo.Session, e *discordgo.MessageCreate, args []string) {
			outputs := make([]string, 0)
			// for _, v := range cmdManager.GetAllCommandKeys() {
			// 	outputs = append(outputs, cmdManager.Get
			// }
			fmt.Println(outputs)
		},
	})
}
