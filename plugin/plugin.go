package plugin

import (
	"os/exec"

	"github.com/hashicorp/go-plugin"
)

var handshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "BASIC_PLUGIN",
	MagicCookieValue: "hello",
}

func NewClient(cmd string) (c *plugin.Client) {
	return plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: handshakeConfig,
		Plugins:         pluginMap,
		Cmd:             exec.Command(cmd, "plugin"),
	})
}

// pluginMap is the map of plugins we can dispense.
var pluginMap = map[string]plugin.Plugin{
	"greeter": &GreeterPlugin{},
}

func Serve(g Greeter) {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: handshakeConfig,
		Plugins:         PluginMap(g),
	})
}

func PluginMap(g Greeter) map[string]plugin.Plugin {
	return map[string]plugin.Plugin{
		"greeter": &GreeterPlugin{PluginFunc: g},
	}
}
