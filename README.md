# MinatsuBot
## Introduction
MinatsuBot is a extensible bot coded in Go.

For the moment it can only connect to Discord, but in the future that might change.

The idea of this bot is to support simple plugins and have a nice PluginAPI to code towards.
As Go doesn't really support any nice ways to load plugins you have to manually load plugins,
but it because of the PluginAPI you can create simple plugins rather quick and then just load the plugin in.

As of Go 1.8. Go now has a Plugin package to load plugins, but I personally don't feel it's that great,
it seems rather over complicated but I might look into it at some point and make the plugins load dynamic.

As this is an early version of the bot, the PluginAPI is bound to have some drastic changes,
but I will try my best to keep things consistent so that any plugins will work throughout the versions.

## Installation
*Todo this*

## Features
This list is meant to list all intended features, implemented or not.
So it acts as a feature and todo list.

#### General
* [x] Load custom plugins
* [ ] Make things more concurrent with Go channels and such

#### Plugin API specific
* [ ] Custom Settings - settings per service
* [ ] Custom Permissions
* [ ] Broadcast - Broadcast between different chat and services
* [ ] Make everything accessible from the Plugin API - Currently you need to access Bot and other packages to do certain things like send a message
* [ ] Support for databases
* [ ] Multi service support - make plugins know what service the event is being called from and support multiple services (slack, discord, custom chat etc.)

#### Plugins
* [ ] Minecraft Chat - read and send messages from/to minecraft (integrate with Logging/stats plugin)
* [ ] Minecraft Status - show status about minecraft servers, player amount, online etc.
* [ ] Xenforo integration - notify about new posts and such, check Xenforo for more fun to do
* [ ] Permission - manage the internal permission API, support integration with different services like Discord
* [x] Info - commands to get information about the bot, uptime, mem etc.
* [x] Help - commands to get information about specific commands.
* [ ] Manage bot - manage the bot from a chat, change settings, stop, restart etc.
* [ ] Admin plugin - kick, ban, auto ban on swear and such
* [ ] Music plugin
* [ ] Minecraft Integration - integrate with various minecraft top plugins to do things like award for achievement and such.
* [ ] Bob ross
* [ ] Custom commands - add custom commands easily (like .website)
* [ ] Anouncements - broadcast messages at specific times or at a set interval
* [ ] Logging/Stats - log from different services and show stats based on the logs
* [ ] AFK - able to set your status to AFK, so whenever someone goes AFK and you get mentioned it will show a message to whoever mentioned you and later on when you come back it will show the notification
* [ ] Cleverbot support
