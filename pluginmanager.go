package minatsubot

type PluginManager struct {
	plugins map[string]*simpleplugin
}

func (p *PluginManager) addPlugin(plugin Plugin, desc PluginDescription) {
	if desc.Name == "" {
		log.Error("Trying to register a plugin with no name.")
		return
	}
	if desc.Version.Get() == "" {
		log.Error("Trying to register a plugin with no version.")
		return
	}

	if p.plugins[desc.Name] != nil {
		log.Errorf("There is already a plugin with the name %s registered.", desc.Name)
		return
	}
	p.plugins[desc.Name] = &simpleplugin{
		enabled:     false,
		plugin:      plugin,
		description: desc,
	}
	log.Infof("Plugin %s has now been added.", desc.Name)
}

func (p *PluginManager) getPlugin(name string) (*simpleplugin, bool) {
	plugin, ok := p.plugins[name]
	return plugin, ok
}

func (p *PluginManager) getPluginDesc(name string) (PluginDescription, bool) {
	if p.plugins[name] == nil {
		return PluginDescription{}, false
	}
	return p.plugins[name].description, true
}

func (p *PluginManager) getAllPlugins() map[string]*simpleplugin {
	return p.plugins
}

func (p *PluginManager) getAllPluginsDesc() []PluginDescription {
	plugins := []PluginDescription{}

	for _, plugin := range p.plugins {
		plugins = append(plugins, plugin.description)
	}
	return plugins
}

func (p *PluginManager) init() {
	// Initialize all Plugins here
	log.Info("Initializing all plugins.")
	for _, plugin := range p.plugins {
		log.Infof("Initializing %s v%s", plugin.GetDescription().Name, plugin.GetDescription().Version.Get())
		plugin.GetPlugin().Init()
		log.Infof("%s v%s initialized", plugin.GetDescription().Name, plugin.GetDescription().Version.Get())
	}
}

func (p *PluginManager) enable() {
	log.Info("Enabling all plugins.")
	for _, plugin := range p.plugins {
		if plugin.IsEnabled() {
			continue
		}

		log.Infof("Enabling %s v%s", plugin.GetDescription().Name, plugin.GetDescription().Version.Get())
		plugin.GetPlugin().Enable()
		plugin.enabled = true
		log.Infof("%s v%s enabled", plugin.GetDescription().Name, plugin.GetDescription().Version.Get())
	}
}

func (p *PluginManager) disable() {
	for _, plugin := range p.plugins {
		if !plugin.IsEnabled() {
			continue
		}
		log.Infof("Disabling %s v%s", plugin.GetDescription().Name, plugin.GetDescription().Version.Get())
		plugin.GetPlugin().Disable()
		log.Infof("%s v%s disabled", plugin.GetDescription().Name, plugin.GetDescription().Version.Get())
	}
}
