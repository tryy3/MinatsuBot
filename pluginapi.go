package minatsubot

import "github.com/bwmarrin/discordgo"
import "time"

type API struct {
	Event          *EventHandler
	Session        *discordgo.Session
	bot            *Bot
	pluginManager  *PluginManager
	commandManager *CommandManager
}

// GetBotID Returns the bot id
func (api *API) GetBotID() string {
	return api.bot.ID
}

// GetSettings Returns the global bot settings
func (api *API) GetSettings() Settings {
	return api.bot.Settings
}

// RegisterCommand Register a plugins command
func (api *API) RegisterCommand(desc CommandDescription, cmd Command) {
	api.commandManager.registerCommand(desc, cmd)
}

// GetPlugin Return a registered plugin
func (api *API) GetPlugin(name string) (*PluginInfo, bool) {
	return api.pluginManager.getPlugin(name)
}

// GetPluginDesc Returns a PluginDescription of a specific plugin
func (api *API) GetPluginDesc(name string) (PluginDescription, bool) {
	return api.pluginManager.getPluginDesc(name)
}

// GetAllPlugins Returns all existing plugins
func (api *API) GetAllPlugins() map[string]*PluginInfo {
	return api.pluginManager.getAllPlugins()
}

// GetAllPluginsDesc Returns a slice of all existing plugins PluginDescription
func (api *API) GetAllPluginsDesc() []PluginDescription {
	return api.pluginManager.getAllPluginsDesc()
}

// GetCommand Returns a registered commands CommandDescription
func (api *API) GetCommand(name string) (CommandDescription, bool) {
	return api.commandManager.getCommand(name)
}

// GetAllCommands Returns a slice of all registered commands CommandDescription
func (api *API) GetAllCommands() []CommandDescription {
	return api.commandManager.getAllCommands()
}

// GetBootTime Returns a time.Time from when the bot was initialized
func (api *API) GetBootTime() time.Time {
	return api.bot.Starttime
}

// GetBotDescription Returns the bot Description
func (api *API) GetBotDescription() Description {
	return api.bot.Description
}

type PluginInfo struct {
	plugin      Plugin
	enabled     bool
	description PluginDescription
}

func (p PluginInfo) GetPlugin() Plugin {
	return p.plugin
}

func (p PluginInfo) IsEnabled() bool {
	return p.enabled
}

func (p PluginInfo) GetDescription() PluginDescription {
	return p.description
}
