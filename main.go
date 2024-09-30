package main

import (
	pluginInterface "github.com/StandardRunbook/plugin-interface/shared"
	"github.com/StandardRunbook/plugin-template-go/pkg/script"
	"github.com/hashicorp/go-plugin"
)

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: pluginInterface.Handshake,
		Plugins: map[string]plugin.Plugin{
			"kv": &pluginInterface.GRPCPlugin{Impl: &script.Template{}},
		},

		// A non-nil value here enables gRPC serving for this script...
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
