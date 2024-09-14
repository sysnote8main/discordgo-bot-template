package commands

import "github.com/bwmarrin/discordgo"

type TextCommand struct {
	CommandName string
	Aliases     []string
	Description string
	HandlerFunc func(s *discordgo.Session, e *discordgo.MessageCreate, args []string) error
}

type SlashCommand struct {
	Command     *discordgo.ApplicationCommand
	HandlerFunc func(s *discordgo.Session, event *discordgo.InteractionCreate, args []string) error
}
