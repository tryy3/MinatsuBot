package main

import (
	"flag"

	"github.com/tryy3/minatsubot"
	"github.com/tryy3/minatsubot/plugins/helpplugin"
	"github.com/tryy3/minatsubot/plugins/infoplugin"
	"github.com/tryy3/minatsubot/plugins/permissionmanager"
)

var (
	token   string
	prefix  string
	logging string
)

func init() {
	flag.StringVar(&token, "token", "", "Discord Token")
	flag.StringVar(&prefix, "prefix", "!", "Bot prefix")
	flag.StringVar(&logging, "logging", "info", "Bot logging level")
	flag.Parse()
}

func main() {
	bot := minatsubot.NewBot(minatsubot.Settings{
		Token:   token,
		Logging: logging,
		Prefix:  prefix,
	})

	bot.AddPlugin(helpplugin.Help{}, minatsubot.PluginDescription{
		Name: "HelpPlugin",
		Description: minatsubot.Description{
			Version: &minatsubot.Version{"0", "0", "1"},
		},
	})

	bot.AddPlugin(infoplugin.Info{}, minatsubot.PluginDescription{
		Name: "InfoPlugin",
		Description: minatsubot.Description{
			Version: &minatsubot.Version{"0", "0", "1"},
		},
	})

	bot.AddPlugin(permissionmanager.Permission{}, minatsubot.PluginDescription{
		Name: "PermissionManager",
		Description: minatsubot.Description{
			Version: &minatsubot.Version{"0", "0", "1"},
		},
	})

	bot.Init()
}
