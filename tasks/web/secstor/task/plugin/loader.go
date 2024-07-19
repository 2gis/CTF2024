package plugin

import (
	"fmt"
	"gopher/utils/files"
	"path/filepath"
	"plugin"
)

func EnableAllPlugins() {
	pluginsDir := files.Directory{Path: "plugins"}
	err := pluginsDir.CreateAll()
	if err != nil {
		return
	}
	all_plugins, err := filepath.Glob("plugins/*.so")
	if err != nil {
		return
	}

	for _, filename := range all_plugins {
		p, err := plugin.Open(filename)
		if err != nil {
			return
		}

		symbolPluginName, err := p.Lookup("Name")
		pluginName, ok := symbolPluginName.(*string)
		if err != nil || !ok {
			return
		}

		symbolPluginVersion, err := p.Lookup("Version")
		pluginVersion, ok := symbolPluginVersion.(*string)
		if err != nil || !ok {
			fmt.Println("Plugin has no \"Version\" const")
		}

		symbolOnEnable, err := p.Lookup("OnEnable")
		onEnable, ok := symbolOnEnable.(func())
		if err != nil || !ok {
			fmt.Println("Plugin has no \"OnEnable()\" function")
		}

		symbolOnDisable, err := p.Lookup("OnDisable")
		onDisable, ok := symbolOnDisable.(func())
		if err != nil || !ok {
			fmt.Println("Plugin has no \"OnDisable()\" function")
		}

		plug := Plugin{
			Name:      *pluginName,
			Version:   *pluginVersion,
			OnEnable:  onEnable,
			OnDisable: onDisable,
			plugin:    p,
		}
		_, exists := pluginStorage[*pluginName]
		if exists {
			pluginStorage[*pluginName][*pluginVersion] = &plug
		} else {
			pluginStorage[*pluginName] = make(map[string]*Plugin)
			pluginStorage[*pluginName][*pluginVersion] = &plug
		}
	}
	pluginsCounter := 0
	for _, vers := range pluginStorage {
		for _, plug := range vers {
			if plug.OnEnable != nil {
				plug.OnEnable()
			}
			pluginsCounter++
		}
	}
}

func DisableAllPlugins() {
	pluginsCounter := 0
	for _, vers := range pluginStorage {
		for _, plug := range vers {
			if plug.OnDisable != nil {
				plug.OnDisable()
			}
			pluginsCounter++
		}
	}
	pluginStorage = make(map[string]map[string]*Plugin)
}