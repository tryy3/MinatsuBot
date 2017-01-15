package minatsubot

import "github.com/bwmarrin/discordgo"

type PermissionHandler interface {
	HasUserPermission(name, perm string) bool
	HasGroupPermission(name, perm string) bool
}

type Plugin interface {
	Enable()
	Init()
	Disable()
}

type Command interface {
	Run(command string, label string, args []string, message *discordgo.MessageCreate)
}
