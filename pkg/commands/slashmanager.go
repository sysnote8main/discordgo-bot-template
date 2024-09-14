package commands

import (
	"log/slog"
	"maps"
	"slices"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type SlashCommandManager struct {
	cmdMap           map[string]SlashCommand
	helpFieldMap     map[string]*discordgo.MessageEmbedField
	registeredCmdIds []string
}

func (m *SlashCommandManager) ExistCommand(cmdName string) bool {
	_, ok := m.cmdMap[cmdName]
	return ok
}

func (m *SlashCommandManager) AddCommand(command SlashCommand) {
	m.cmdMap[command.Command.Name] = command
	// TODO add to help field map
}

func (m *SlashCommandManager) RemoveCommand(cmdName string) {
	delete(m.cmdMap, cmdName)
	// TODO remove from help field map
}

func (m *SlashCommandManager) GetAllCommandKeys() []string {
	return slices.Sorted(maps.Keys(m.cmdMap))
}

func (m *SlashCommandManager) GetAllCommands() []SlashCommand {
	return slices.SortedFunc(maps.Values(m.cmdMap), func(a, b SlashCommand) int {
		return strings.Compare(a.Command.Name, b.Command.Name)
	})
}

func (m *SlashCommandManager) GetAllRegisteredCommandKeys() []string {
	return m.registeredCmdIds
}

// TODO add help field getter

func (m *SlashCommandManager) OnEvent(session *discordgo.Session, event *discordgo.InteractionCreate) {
	v, ok := m.cmdMap[event.ApplicationCommandData().Name]
	if ok {
		err := v.HandlerFunc(session, event)
		if err != nil {
			slog.Error("Failed to process slash command", slog.String("cmdName", v.Command.Name), slog.Any("error", err))
		}
	}
}

func (m *SlashCommandManager) Register(session *discordgo.Session, guildId string) {
	if guildId == "" {
		slog.Warn("GuildId is empty.")
		// TODO remove this on production
		// On production, needs to register for all guilds & users
		return
	}
	slog.Info("Registering slash commands...", slog.Int("allCmdCount", len(m.cmdMap)), slog.String("guildId", guildId))
	for k, v := range m.cmdMap {
		cmd, err := session.ApplicationCommandCreate(session.State.User.ID, guildId, v.Command)
		if err != nil {
			slog.Error("Failed to create slash command", slog.String("cmdId", k))
		} else {
			slog.Debug("Command was registered!", slog.String("internalId", k), slog.String("discordSideId", cmd.ID))
			m.registeredCmdIds = append(m.registeredCmdIds, cmd.ID)
		}
	}
	slog.Info("Finished to register slash commands.", slog.Int("allCmdCount", len(m.cmdMap)), slog.Int("registeredCmdCount", len(m.registeredCmdIds)))
}

func NewSlashCommandManager() *SlashCommandManager {
	return &SlashCommandManager{
		cmdMap:           map[string]SlashCommand{},
		helpFieldMap:     map[string]*discordgo.MessageEmbedField{},
		registeredCmdIds: make([]string, 0),
	}
}
