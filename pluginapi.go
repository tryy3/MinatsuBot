package minatsubot

import "github.com/bwmarrin/discordgo"
import "time"

type API struct {
	Event          *EventHandler
	Session        *discordgo.Session
	bot            *Bot
	pluginManager  *PluginManager
	commandManager *CommandManager
	permission     *permissionManager
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
func (api *API) GetPlugin(name string) (*simpleplugin, bool) {
	return api.pluginManager.getPlugin(name)
}

// GetPluginDesc Returns a PluginDescription of a specific plugin
func (api *API) GetPluginDesc(name string) (PluginDescription, bool) {
	return api.pluginManager.getPluginDesc(name)
}

// GetAllPlugins Returns all existing plugins
func (api *API) GetAllPlugins() map[string]*simpleplugin {
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

// GetLogger Returns a new Logger
func (api *API) GetLogger(name string) *Logger {
	return newLogger(name, getLoggingLevel(api.GetSettings().Logging))
}

// HasUserPermission Checks if the user has access to a certain permission.
// If no PermissionHandler has been assigned, this will always return true.
func (api *API) HasUserPermission(name, perm string) bool {
	return api.permission.hasUserPermission(name, perm)
}

// HasGroupPermission Checks if the group has access to a certain permission.
// If no PermissionHandler has been assigned, this will always return true.
func (api *API) HasGroupPermission(name, perm string) bool {
	return api.permission.hasGroupPermission(name, perm)
}

// SetPermissionHandler Sets the bots PermissionHandler to handler.
// Can only contain one PermissionHandler, so the return bool means
// if it was able to set a PermissionHandler or not
func (api *API) SetPermissionHandler(handler PermissionHandler) bool {
	return api.permission.setPermissionHandler(handler)
}

// GetPermissionHandler Returns the bot PermissionHandler.
func (api *API) GetPermissionHandler() PermissionHandler {
	return api.permission.getManager()
}
