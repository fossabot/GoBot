package plugin

type GoBotPlugin interface {
	LoadPlugin() error
	UnloadPlugin() error
	PluginInfo() PluginInfo
}
