package minatsubot

import (
	"sync"
	"time"

	"os"

	"io"

	"github.com/bwmarrin/discordgo"
)

var (
	Log       Logger
	PluginAPI *API
)

func init() {
	PluginAPI = &API{
		Event: &EventHandler{
			mu:  &sync.Mutex{},
			rmu: &sync.RWMutex{},
		},
	}
	Log = Logger{
		Writers: []io.Writer{os.Stdout},
		Prefix:  "MinatsuBot",
		Level:   InfoLevel,
	}
}

func NewBot(settings Settings) *Bot {
	// TODO: Make the settings get auto loaded
	bot := &Bot{
		Settings: settings,
		Description: Description{
			Version: &Version{"0", "0", "1"},
			Author:  "tryy3",
			Website: "google.com",
			Info:    "Extensible Chat Bot",
		},
	}

	Log.Level = GetLoggingLevel(bot.Settings.Logging)
	PluginAPI.bot = bot
	PluginAPI.commandManager = &CommandManager{commands: []simplecommand{}}
	PluginAPI.pluginManager = &PluginManager{plugins: map[string]*PluginInfo{}}

	return bot
}

type Bot struct {
	Settings    Settings
	Starttime   time.Time
	Description Description
	ID          string
}

func (b *Bot) AddPlugin(plugin Plugin, desc PluginDescription) {
	PluginAPI.pluginManager.addPlugin(plugin, desc)
}

func (b *Bot) Init() {
	Log.Info("Initializing the Bot")
	b.Starttime = time.Now()

	Log.Info("Initializing the CommandManager")
	PluginAPI.Event.AddHandler(PluginAPI.commandManager.handler)

	Log.Info("Initializing the PluginManager")
	PluginAPI.pluginManager.init()

	Log.Info("Creating a discord session")
	discord, err := discordgo.New(b.Settings.Token)
	if err != nil {
		Log.Error("Error creating a Discord session,", err)
		PluginAPI.pluginManager.disable()
		return
	}

	Log.Info("Gathering bot data")
	user, err := discord.User("@me")
	if err != nil {
		Log.Error("Error obtaining account details,", err)
		PluginAPI.pluginManager.disable()
		return
	}

	// Save the bots ID for later use
	b.ID = user.ID

	// Add a discord chat handler
	discord.AddHandler(PluginAPI.Event.handler)

	Log.Info("Opening discord connection")
	err = discord.Open()
	if err != nil {
		Log.Error("Error oepning discord connection,", err)
		PluginAPI.pluginManager.disable()
		return
	}

	PluginAPI.Session = discord
	PluginAPI.pluginManager.enable()

	Log.Info("Bot is now running. Press CTRL-C to exit.")
	<-make(chan struct{})
	return
}
