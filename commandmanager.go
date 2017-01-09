package minatsubot

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

type simplecommand struct {
	description CommandDescription
	cmd         Command
}

type CommandManager struct {
	commands []simplecommand
}

func (c *CommandManager) registerCommand(desc CommandDescription, cmd Command) {
	c.commands = append(c.commands, simplecommand{desc, cmd})
}

func (c *CommandManager) findCommand(name string) (simplecommand, string, bool) {
	n := strings.ToLower(name)
	for _, scmd := range c.commands {
		if strings.ToLower(scmd.description.Name) == n {
			return scmd, name, true
		}
		for _, alias := range scmd.description.Aliases {
			if strings.ToLower(alias) == n {
				return scmd, alias, true
			}
		}
	}
	return simplecommand{}, "", false
}

func (c *CommandManager) getCommand(name string) (CommandDescription, bool) {
	cmd, _, ok := c.findCommand(name)
	if !ok {
		return CommandDescription{}, false
	}
	return cmd.description, true
}

func (c *CommandManager) getAllCommands() []CommandDescription {
	cmds := []CommandDescription{}
	for _, cmd := range c.commands {
		cmds = append(cmds, cmd.description)
	}
	return cmds
}

func (c *CommandManager) getPrefix(message string) (string, []string) {
	args := strings.Split(message, " ")
	if strings.HasPrefix(message, PluginAPI.GetSettings().Prefix) {
		name := strings.TrimPrefix(args[0], PluginAPI.GetSettings().Prefix)
		return name, args[1:]
	}

	if args[0] == ("<@!" + PluginAPI.GetBotID() + ">") {
		if len(args) <= 1 {
			return "", nil
		}

		return args[1], args[2:]
	}
	return "", nil
}

func (c *CommandManager) handler(m *discordgo.MessageCreate) {
	if m.Author.ID == PluginAPI.GetBotID() {
		return
	}

	commandName, args := c.getPrefix(m.Content)

	if commandName == "" {
		return
	}

	command, label, ok := c.findCommand(commandName)

	if !ok {
		return
	}
	command.cmd.Run(commandName, label, args, m)
}
