package commands

import (
	"log/slog"
	"maps"
	"slices"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type TextCommandManager struct {
	commandMap      map[string]TextCommand
	prefix          string
	helpCmdFieldMap map[string]*discordgo.MessageEmbedField
}

func (m *TextCommandManager) ExistCommand(cmdName string) bool {
	_, ok := m.commandMap[cmdName]
	return ok
}

func (m *TextCommandManager) AddCommand(cmd TextCommand) {
	m.commandMap[cmd.CommandName] = cmd
	m.helpCmdFieldMap[cmd.CommandName] = makeFieldFromTextCommand(cmd)
}

func (m *TextCommandManager) RemoveCommand(cmdName string) {
	delete(m.commandMap, cmdName)
	delete(m.helpCmdFieldMap, cmdName)
}

func (m *TextCommandManager) GetAllCommandKeys() []string {
	return slices.Sorted(maps.Keys(m.commandMap))
}

func (m *TextCommandManager) GetAllCommands() []TextCommand {
	return slices.SortedFunc(maps.Values(m.commandMap), func(a, b TextCommand) int {
		return strings.Compare(a.CommandName, b.CommandName)
	})
}

func (m *TextCommandManager) GetHelpCommandFields() []*discordgo.MessageEmbedField {
	return slices.SortedFunc(maps.Values(m.helpCmdFieldMap), func(a, b *discordgo.MessageEmbedField) int {
		return strings.Compare(a.Name, b.Name)
	})
}

func (m *TextCommandManager) OnEvent(session *discordgo.Session, event *discordgo.MessageCreate) {
	trimStr := strings.TrimPrefix(event.Message.Content, m.prefix)
	if trimStr == event.Message.Content {
		// Prefix is not found in message
		return
	}
	args := strings.Split(trimStr, " ")
	cmdName := args[0]
	v, ok := m.commandMap[cmdName]
	if ok {
		err := v.HandlerFunc(session, event, args[1:])
		if err != nil {
			session.ChannelMessageSend(event.ChannelID, "Failed to process command.")
			slog.Error("Failed to process command", slog.String("cmdName", cmdName), slog.Any("error", err))
		}
	}
}

func makeFieldFromTextCommand(cmd TextCommand) *discordgo.MessageEmbedField {
	return &discordgo.MessageEmbedField{
		Name:  cmd.CommandName,
		Value: cmd.Description,
	}
}

func NewTextCommandManager(prefix string) *TextCommandManager {
	return &TextCommandManager{
		commandMap:      map[string]TextCommand{},
		helpCmdFieldMap: map[string]*discordgo.MessageEmbedField{},
		prefix:          prefix,
	}
}
