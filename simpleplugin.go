package minatsubot

type simpleplugin struct {
	plugin      Plugin
	enabled     bool
	description PluginDescription
}

func (p simpleplugin) GetPlugin() Plugin {
	return p.plugin
}

func (p simpleplugin) IsEnabled() bool {
	return p.enabled
}

func (p simpleplugin) GetDescription() PluginDescription {
	return p.description
}
