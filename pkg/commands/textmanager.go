package commands

import (
	"maps"
	"slices"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var (
	prefix = "!"
)

type TextCommandManager struct {
	commandMap map[string]TextCommand
}

func (m *TextCommandManager) ExistCommand(cmdName string) bool {
	_, ok := m.commandMap[cmdName]
	return ok
}

func (m *TextCommandManager) AddCommand(cmd TextCommand) {
	m.commandMap[cmd.CommandName] = cmd
}

func (m *TextCommandManager) RemoveCommand(cmdName string) {
	delete(m.commandMap, cmdName)
}

func (m *TextCommandManager) GetAllCommandKeys() []string {
	return slices.Sorted(maps.Keys(m.commandMap))
}

func (m *TextCommandManager) GetAllCommands() []TextCommand {
	return slices.SortedFunc(maps.Values(m.commandMap), func(a, b TextCommand) int {
		return strings.Compare(a.CommandName, b.CommandName)
	})
}

func (m *TextCommandManager) OnEvent(session *discordgo.Session, event *discordgo.MessageCreate) {
	trimStr := strings.TrimPrefix(event.Message.Content, prefix)
	if trimStr == event.Message.Content {
		// Prefix is not found in message
		return
	}
	args := strings.Split(trimStr, " ")
	cmdName := args[0]
	v, ok := m.commandMap[cmdName]
	if ok {
		v.HandlerFunc(session, event, args[1:])
	}
}

func NewTextCommandManager() *TextCommandManager {
	return &TextCommandManager{
		commandMap: map[string]TextCommand{},
	}
}
