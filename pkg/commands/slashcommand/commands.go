package slashcommand

import (
	"github.com/bwmarrin/discordgo"

	"github.com/sysnote8main/discordgo-bot-template/pkg/commands"
)

func Register(cmdManager *commands.SlashCommandManager) {
	cmdManager.AddCommand(commands.SlashCommand{
		Command: &discordgo.ApplicationCommand{
			Name:        "ping",
			Description: "pong!",
		},
		HandlerFunc: func(s *discordgo.Session, event *discordgo.InteractionCreate) error {
			return s.InteractionRespond(event.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "pong!",
					Flags:   discordgo.MessageFlagsEphemeral,
				},
			})
		},
	})
}