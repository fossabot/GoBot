package plugin

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"plugin"
	"runtime"
	"sync"

	"github.com/IcedTechLab/GoBot/platform"
)

type PluginManager struct {
	plugins []Plugin
}

type PluginInfo struct {
	Author  string
	Version string
	About   string
	Name    string
}

type Plugin struct {
	Info   PluginInfo
	Handle GoBotPlugin
}

var manager_ins PluginManager
var manager_once sync.Once

func NewManager() PluginManager {
	manager_once.Do(func() {})
	return manager_ins
}

func (p PluginManager) GetPlugins() []Plugin {
	return p.plugins
}

func (p *PluginManager) LoadPlugin() {

	if runtime.GOOS == "windows" {
		LOG_UNSUPPORTED_OS()
		return
	}

	wd, _ := os.Getwd()
	pluginDir := filepath.Join(wd, "plugins")
	if dirs, err := ioutil.ReadDir(pluginDir); err != nil {
		return
	} else {
		fileList := make([]string, 0)
		// Search plugin file
		for _, v := range dirs {
			if v.IsDir() {
				continue
			}
			if pluginFilePath := filepath.Join(pluginDir, v.Name()); filepath.Ext(pluginFilePath) == platform.PLUGIN_EXTENSION {
				fileList = append(fileList, pluginFilePath)
			}
		}
		// Load plugin file
		for _, v := range fileList {
			if pluginHandle, err := plugin.Open(v); err != nil {
				continue
			} else {
				// Lookup symbol
				if sym, err := pluginHandle.Lookup("LoadPlugin"); err != nil {
					continue
				} else {
					// Assert
					plugin, ok := sym.(GoBotPlugin)
					if ok {
						p.plugins = append(p.plugins, Plugin{Info: plugin.PluginInfo(), Handle: plugin})
					}
				}
			}
		}
		for _, v := range p.plugins {
			if err := v.Handle.LoadPlugin(); err != nil {
				LOG_PLUGIN_LOAD_STATUS(v.Info.Name, err, true)
			} else {
				LOG_PLUGIN_LOAD_STATUS(v.Info.Name, err, false)
			}
		}
	}
}

func (p *PluginManager) UnloadPlugin() {
	for _, v := range p.plugins {
		if err := v.Handle.UnloadPlugin(); err != nil {
			LOG_PLUGIN_UNLOAD_STATUS(v.Info.Name, err, true)
		} else {
			LOG_PLUGIN_UNLOAD_STATUS(v.Info.Name, err, false)
		}
	}
}
