package plugin

import (
	"plugin"
)

var pluginStorage = make(map[string]map[string]*Plugin)

type Plugin struct {
	Name      string
	Version   string
	OnEnable  func()
	OnDisable func()
	plugin    *plugin.Plugin
}

func GetPlugin(name, version string) (*Plugin, bool) {
	vers, exists_vers := pluginStorage[name]
	if !exists_vers {
		return nil, false
	}
	if version == "" {
		plug, exists := vers[""]
		if exists {
			return plug, exists
		} else {
			max := "0.0.0.0.0.0.0"
			for version_i, _ := range vers {
				i := 0
				for _, char := range version_i {
					if char > []rune(max)[i] {
						max = version_i
						break
					}
					i++
				}
			}
			if max == "" {
				return nil, false
			} else {
				return vers[max], true
			}
		}
	}
	plug, exists := vers[version]
	return plug, exists
}

func (plugin *Plugin) Lookup(name string) (plugin.Symbol, error) {
	return plugin.plugin.Lookup(name)
}