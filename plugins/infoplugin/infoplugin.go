package infoplugin

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/dustin/go-humanize"
	"github.com/shirou/gopsutil/mem"
	"github.com/tryy3/minatsubot"
)

var (
	Log minatsubot.Logger
)

type Info struct {
}

func (i Info) Init() {
	Log = minatsubot.Logger{
		Writers: []io.Writer{os.Stdout},
		Prefix:  "InfoPlugin",
		Level:   minatsubot.GetLoggingLevel(minatsubot.PluginAPI.GetSettings().Logging),
	}
}

func (i Info) Enable() {
	minatsubot.PluginAPI.RegisterCommand(
		minatsubot.CommandDescription{
			Name:        "info",
			Description: "Information about the bot",
			Usage:       "<command>",
			Aliases:     []string{"stats"},
		}, &InfoCMD{})
}

func (i Info) Disable() {

}

type InfoCMD struct {
}

func (i InfoCMD) Run(command string, label string, args []string, message *discordgo.MessageCreate) {
	stats := runtime.MemStats{}
	runtime.ReadMemStats(&stats)

	sys, err := mem.VirtualMemory()
	sysMem := ""
	if err == nil {
		sysMem = fmt.Sprintf("%s (%.0f%%), %s",
			humanize.Bytes(sys.Used),
			sys.UsedPercent,
			humanize.Bytes(sys.Total),
		)
	} else {
		sysMem = "Failed collecting memory information"
		Log.Warn("Failed collecting memory information")
	}

	description := minatsubot.PluginAPI.GetBotDescription()
	procMem := fmt.Sprintf("%s / %s (%s garbage collected)",
		humanize.Bytes(stats.Alloc),
		humanize.Bytes(stats.Sys),
		humanize.Bytes(stats.TotalAlloc),
	)
	pluginsEnabled := 0
	pluginsFound := len(minatsubot.PluginAPI.GetAllPlugins())
	commands := len(minatsubot.PluginAPI.GetAllCommands())

	for _, plugin := range minatsubot.PluginAPI.GetAllPlugins() {
		if plugin.IsEnabled() {
			pluginsEnabled++
		}
	}

	user := minatsubot.PluginAPI.Session.State.User

	minatsubot.PluginAPI.Session.ChannelMessageSendEmbed(message.ChannelID, &discordgo.MessageEmbed{
		Title:       "Information",
		Description: description.Info,
		Color:       16726485,
		Author: &discordgo.MessageEmbedAuthor{
			Name:    user.Username,
			IconURL: discordgo.EndpointUserAvatar(minatsubot.PluginAPI.GetBotID(), user.Avatar),
		},
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{Name: "Version", Value: description.Version.Get(), Inline: true},
			&discordgo.MessageEmbedField{Name: "DiscordGo", Value: discordgo.VERSION, Inline: true},
			&discordgo.MessageEmbedField{Name: "Go", Value: runtime.Version(), Inline: true},
			&discordgo.MessageEmbedField{Name: "Uptime", Value: getDuration(time.Since(minatsubot.PluginAPI.GetBootTime())), Inline: true},
			&discordgo.MessageEmbedField{Name: "Bot memory (alloc, sys, freed)", Value: procMem, Inline: true},
			&discordgo.MessageEmbedField{Name: "System memory (used, total)", Value: sysMem, Inline: true},
			&discordgo.MessageEmbedField{Name: "Plugins", Value: fmt.Sprintf("%d (%d enabled)", pluginsFound, pluginsEnabled), Inline: true},
			&discordgo.MessageEmbedField{Name: "Commands", Value: strconv.Itoa(commands), Inline: true},
			&discordgo.MessageEmbedField{Name: "Website", Value: description.Website, Inline: true},
			&discordgo.MessageEmbedField{Name: "Author", Value: description.Author, Inline: true},
		},
	})
}

func getDuration(duration time.Duration) string {
	return fmt.Sprintf("%0.2d:%02d:%02d",
		int(duration.Hours()),
		int(duration.Minutes())%60,
		int(duration.Seconds())%60,
	)
}
